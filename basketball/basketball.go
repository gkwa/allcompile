package basketball

type BasketBall struct {
	Brand  string
	Color  string
	Weight int
}

type basketballFactory struct {
	Brand  string
	Color  string
	Weight int
}

const (
	DefaultBrand  = "DefaultBasketballBrand"
	DefaultColor  = "DefaultBasketballColor"
	DefaultWeight = 999
)

type BallFactory interface {
	CreateBasketBall() *BasketBall
}

func NewBasketballFactory(options ...func(*basketballFactory)) BallFactory {
	factory := &basketballFactory{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	for _, option := range options {
		option(factory)
	}

	return factory
}

func (f *basketballFactory) CreateBasketBall() *BasketBall {
	ball := &BasketBall{
		Brand:  f.Brand,
		Color:  f.Color,
		Weight: f.Weight,
	}

	return ball
}

func WithBrand(brand string) func(*basketballFactory) {
	return func(f *basketballFactory) {
		f.Brand = brand
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
