package main

import (
	"flag"
	"fmt"
	"fs/gcloud"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"utils/markov"
)

var fInit = flag.Bool("init", false, "do you want to re-init your model?")
var fDir = flag.String("m", "./markov-model", "where to store the model?")
var fDepth = flag.Int("d", 50, "how many tokens do you want the next token based on?")
var fCorpus = flag.String("f", "", "the text file you want to work with")
var fCount = flag.Int("reps", 10, "number of passages to generate")
var fStart = flag.String("s", "I have", "what do you want to seed the text with?")
var fLength = flag.Int("l", 240, "how long do you want the output to be?")

func main() {
	flag.Parse()

	store, err := markov.NewDataStore(gcloud.Env().ProjectId(), "testing")
	if err != nil {
		log.Fatalf("err opening ds\n%s", err)
	}
	model := markov.NewModelBuilder(store, 3)

	corpus, err := os.Open(*fCorpus)
	if err != nil {
		log.Fatalf("err opening corpus\n%s", err)
	}

	if err := markov.Ingest(model, corpus, markov.NewWordScanner(markov.DefaultWordEquivalentChars, markov.DefaultExcludedChars), *fDepth); err != nil {
		log.Fatalf("err ingesting corpus\n%s", err)
	}

	if errs := model.Close(); errs != nil {
		log.Fatalf("err ingesting corpus\n%v", errs)
	}

	chooser := markov.NewRandomChooser()
	execer := markov.NewModelExecer(store)
	for r := 0; r < *fCount; r++ {
		b, err := markov.Exec(execer, chooser, strings.Split(*fStart, " "), *fLength, printFunc, *fLength)
		if err != nil {
			log.Fatalf("%s", err)
		}
		out, _ := ioutil.ReadAll(b)
		fmt.Println(string(out))
	}

}

func printFunc(token string, pos int) string {
	if len(token) == 0 {
		return token
	}
	switch token[0] {
	case '"', '\'', '.', ';', ',', ':':
		return token
	default:
		return " " + token
	}
}
