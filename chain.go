package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/rand"
)

type Chain struct {
	chainLength int
	data        map[string][]string
}

func NewChain(chainLength int) *Chain {
	return &Chain{
		chainLength: chainLength,
		data:        make(map[string][]string),
	}
}

func (c *Chain) Build(words []string) {
	for i := 0; i < len(words)-c.chainLength; i++ {
		key := strings.Join(words[i:i+c.chainLength], " ")
		value := words[i+c.chainLength]
		c.data[key] = append(c.data[key], value)
	}
	return
}

func (c *Chain) Generate(outputLength int) string {
	keys := []string{}
	for k := range c.data {
		keys = append(keys, k)
	}

	output := strings.Split(keys[rand.Intn(len(keys))], " ")

	for i := 0; i < outputLength; i++ {
		key := strings.Join(output[i:i+c.chainLength], " ")
		words := c.data[key]
		if len(words) == 0 {
			break
		}
		output = append(output, words[rand.Intn(len(words))])
	}

	return strings.Join(output, " ")
}

func (c *Chain) Print(outputLength int) {
	fmt.Println(c.Generate(outputLength))
}
