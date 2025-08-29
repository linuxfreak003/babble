package markov

type Generator[State any] interface {
	Initialize()
	Generate(State) State
}

type Chain[State any] struct {
	current State
	gen     Generator[State]
}

func NewChain[State any](g Generator[State]) *Chain[State] {
	g.Initialize()
	return &Chain[State]{
		gen: g,
	}
}

func (c *Chain[State]) NextState() State {
	c.current = c.gen.Generate(c.current)
	return c.current
}
