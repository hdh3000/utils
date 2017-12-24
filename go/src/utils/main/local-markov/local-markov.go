package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"bytes"
	"fmt"
	"strings"
	"utils/markov"
	"utils/markov/local"
)

// Required!
var fDepth = flag.Int("d", 10, "how many tokens do you want the next token based on")
var fModel = flag.String("m", "", "where is the json model?")

// For Init
var fCorpus = flag.String("f", "", "the text file you want to work with")

// For Gen
var fLength = flag.Int("l", 240, "how long do you want the output to be?")
var fStart = flag.String("s", "I have", "what do you want to seed the text with?")

var fCount = flag.Int("reps", 10, "number of passages to generate")

func main() {
	flag.Parse()

	model, err := local.NewModel("")
	if err != nil {
		log.Fatalf("failed to get model\n%s", err)
	}

	f, err := os.Open(*fCorpus)
	if err != nil {
		log.Fatalf("err opening dir\n%s", err)
	}

	s := bufio.NewScanner(f)
	s.Split(local.Scan)

	var tokens []string
	for s.Scan() {
		text := s.Text()
		text = strings.Map(mapper, text)

		if len(tokens) < *fDepth+1 {
			tokens = append(tokens, text)
			continue
		}

		model.Add(tokens)
		tokens = tokens[1:]
		tokens = append(tokens, text)

	}

	chooser := markov.NewRandomChooser()

	for r := 0; r < *fCount; r++ {

		buf := bytes.Buffer{}
		previous := strings.Split(*fStart, " ")
		for i := len(previous); i < *fDepth; i++ {
			choices := model.GetChoices(previous)
			next := chooser.Choose(choices)
			previous = append(previous, next)
		}

		startWriting := false
		for i := 0; i < *fLength; i++ {
			choices := model.GetChoices(previous)
			next := chooser.Choose(choices)
			previous = append(previous[1:], next)


			if startWriting {
				if !local.isPunctuation(next) {
					fmt.Println(" ")
				}
				buf.WriteString(next)
			}

			if next == "." && !startWriting {
				startWriting = true
			}

		}
		fmt.Println(buf.String())
	}
}

func mapper(r rune) rune {
	return rune(strings.ToLower(string(r))[0])
}
