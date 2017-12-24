package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"utils/markov"
)

var fInit = flag.Bool("init", false, "do you want to re-init your model?")
var fDir = flag.String("m", "./markov-model", "where to store the model?")
var fDepth = flag.Int("d", 5, "how many tokens do you want the next token based on?")
var fCorpus = flag.String("f", "", "the text file you want to work with")
var fCount = flag.Int("reps", 10, "number of passages to generate")
var fStart = flag.String("s", "I have", "what do you want to seed the text with?")
var fLength = flag.Int("l", 240, "how long do you want the output to be?")

func main() {
	flag.Parse()

	store, err := markov.NewFSStore(*fDir)
	if err != nil {
		log.Fatalf("%s", err)
	}

	model := markov.NewModelBuilder(store, 20)
	if *fInit {
		if err := os.RemoveAll(*fDir); err != nil {
			log.Fatalf("%s", err)
		}

		if err := os.MkdirAll(*fDir, 0740); err != nil {
			log.Fatalf("%s", err)
		}

		corpus, err := os.Open(*fCorpus)
		if err != nil {
			log.Fatalf("err opening corpus\n%s", err)
		}

		s := bufio.NewScanner(corpus)
		s.Split(markov.NewWordScanner(markov.DefaultWordEquivalentChars, markov.DefaultExcludedChars))

		var tokens []string
		for s.Scan() {
			text := strings.ToLower(s.Text())

			if len(tokens) < *fDepth+1 {
				tokens = append(tokens, text)
				continue
			}

			if err := model.Add(tokens); err != nil {
				log.Fatalf("%s", err)
			}

			tokens = tokens[1:]
			tokens = append(tokens, text)

			if s.Err() != nil {
				log.Fatalf("%s", err)
			}
		}
		model.Close()

		fmt.Println("Model created at %s", *fDir)
		return
	}

	chooser := markov.NewRandomChooser()

	execer := markov.NewModelExecer(store)
	for r := 0; r < *fCount; r++ {
		previous := strings.Split(*fStart, " ")
		for i := len(previous); i < *fDepth; i++ {
			choices, err := execer.GetChoices(previous)
			if err != nil {
				log.Fatalf("%s", err)
			}
			next := chooser.Choose(choices)
			previous = append(previous, next)
		}

		buf := bytes.Buffer{}
		for i := 0; i < *fLength; i++ {
			choices, err := execer.GetChoices(previous)
			if err != nil {
				log.Fatalf("%s", err)
			}
			next := chooser.Choose(choices)
			previous = append(previous[1:], next)
			buf.WriteString(next)

			buf.WriteString(" ")
			if i%15 == 0 && i != 0 {
				buf.WriteString("\n")
			}
		}
		fmt.Println(buf.String())
		fmt.Println("")

	}

}
