package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func DownloadURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error getting url: %s", err)
	}
	defer resp.Body.Close()

	buffer := bytes.NewBuffer(nil)

	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		return "", fmt.Errorf("error copying response body: %s", err)
	}

	return string(buffer.Bytes()), nil
}

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	return string(data), nil
}

func main() {
	var source, sourceType string
	var chainLength, outputLength int

	flag.StringVar(&source, "source", "", "Source used to build markov chain")
	flag.StringVar(&sourceType, "source-type", "file", "Type of source. Options: file, url")
	flag.IntVar(&outputLength, "output", 5, "Output length in number of words")
	flag.IntVar(&chainLength, "chain", 2, "Chain length. A longer change will have less variation, but make more sense")

	flag.Parse()

	var data string
	var err error

	switch sourceType {
	case "file":
		data, err = ReadFile(source)
		if err != nil {
			logrus.Errorf("could not read file: %v", err)
			os.Exit(1)
		}
	case "url":
		data, err = DownloadURL(source)
		if err != nil {
			logrus.Errorf("could not download url: %v", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown source type")
	}

	words := strings.Split(data, " ")

	chain := NewChain(chainLength)
	chain.Build(words)
	chain.Print(outputLength)
}
