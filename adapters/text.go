package adapters

import (
	"math/rand"
	"strings"

	"github.com/linuxfreak003/babble/markov"
)

type textState struct {
	words []string
}

func (s textState) Key() string {
	return strings.Join(s.words, " ")
}

func (s textState) String() string {
	if len(s.words) == 0 {
		return ""
	}
	return s.words[len(s.words)-1]
}

type textGen struct {
	sources []string
	table   map[string][]string
	n       int
}

func NewTextGenerator(n int, sources ...string) markov.Generator[textState] {
	return &textGen{
		sources: sources,
		n:       n,
	}
}

func (g *textGen) Initialize() {
	// Parse source
	g.table = make(map[string][]string)
	for _, source := range g.sources {
		words := strings.Fields(source)
		for i, word := range words {
			var key string
			switch {
			case i == 0:
				// pass
			case i == 1:
				key = words[0]
			case i < g.n:
				key = strings.Join(words[0:i], " ")
			default:
				key = strings.Join(words[i-g.n:i], " ")
			}
			g.table[key] = append(g.table[key], word)
		}
	}
}

func (g *textGen) Generate(s textState) textState {
	words := g.table[s.Key()]
	s.words = append(s.words, GetRandomWord(words))
	if len(s.words) > g.n {
		s.words = s.words[1:]
	}
	return s
}

func GetRandomWord(words []string) string {
	if len(words) == 0 {
		return ""
	}
	return words[rand.Intn(len(words))]
}
