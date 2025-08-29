package adapters

import (
	"fmt"
	"math/rand"

	"github.com/linuxfreak003/babble/markov"
)

type State struct {
	a int
	b int
}

func (s State) String() string {
	return fmt.Sprintf("(%d, %d)", s.a, s.b)
}
func NewNumberGenerator() markov.Generator[State] {
	return &numGen{}
}

type numGen struct{}

func (*numGen) Initialize() {}

func (*numGen) Generate(s State) State {
	if s.a == 0 && s.b == 0 {
		return State{
			a: 0,
			b: 1,
		}
	}
	next := State{
		a: s.b,
	}
	if rand.Intn(2) == 1 {
		next.b = s.a + s.b
	} else {
		next.b = s.b - s.a
		if next.b < 0 {
			next.b *= -1
		}
	}
	return next
}
