package basketball

type BasketBall struct {
	Size   int
	Color  string
	Weight int
}

type BallFactory interface {
	CreateBasketBall() *BasketBall
}

func NewBasketballFactory(options ...func(*basketballFactory)) BallFactory {
	factory := &basketballFactory{
		Size:   7,        // Default values
		Color:  "Orange", // Default values
		Weight: 22,       // Default values
	}

	for _, option := range options {
		option(factory)
	}

	return factory
}

type basketballFactory struct {
	Size   int
	Color  string
	Weight int
}

func (f *basketballFactory) CreateBasketBall() *BasketBall {
	ball := &BasketBall{
		Size:   f.Size,
		Color:  f.Color,
		Weight: f.Weight,
	}

	return ball
}

func WithSize(size int) func(*basketballFactory) {
	return func(f *basketballFactory) {
		f.Size = size
	}
}

func WithColor(color string) func(*basketballFactory) {
	return func(f *basketballFactory) {
		f.Color = color
	}
}

func WithWeight(weight int) func(*basketballFactory) {
	return func(f *basketballFactory) {
		f.Weight = weight
	}
}
