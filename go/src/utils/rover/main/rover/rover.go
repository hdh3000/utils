package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"utils/rover"
)

var recursive = flag.Bool("r", false, "Use the recursive method to find the best combination, rather than the default sequential impl")
var numWorkers = flag.Int("w", 20, "Set the number of workers that will process the data in parallel")

func main() {
	flag.Parse()
	imd, err := unmarshalChunks(os.Stdin)
	if err != nil {
		log.Fatalf("failed to marshal stdin into chunks...\n%s\n", err)
	}

	if *recursive {
		fmt.Println(rover.FindSmallestChunkSetRecursively(imd).String())
	} else {
		if *numWorkers < 0 {
			log.Fatal("-w can't be less than zero")
		}
		fmt.Println(rover.FindSmallestChunkSetSequentially(imd, *numWorkers).String())
	}

}

// unmarshalChunks handles reading chunks from stdin
func unmarshalChunks(buf io.Reader) (*rover.ImgMetadata, error) {
	scanner := bufio.NewScanner(buf)

	imd := &rover.ImgMetadata{}

	line := 0
	for scanner.Scan() {
		if line == 0 {
			size, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}
			// Set the image size, and move on
			imd.Size = size
			line++
			continue
		}

		row := strings.Split(scanner.Text(), "\t")
		if len(row) != 3 {
			return nil, fmt.Errorf("failed to read line %d... %q is not a 3 value tab seperated string", line, scanner.Text())
		}

		start, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, err
		}

		len, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, err
		}

		chunk := rover.Chunk{
			Id:    row[0],
			Start: start,
			Len:   len,
		}

		imd.Chunks = append(imd.Chunks, chunk)
		line++
	}

	return imd, nil
}
