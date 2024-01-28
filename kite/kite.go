package kite

const (
	DefaultBrand  = "DefaultKiteBrand"
	DefaultColor  = "DefaultKiteColor"
	DefaultWeight = 16
)

type Kite struct {
	Brand  string
	Color  string
	Weight int
}

type KiteOptions struct {
	Brand  string
	Color  string
	Weight int
}

type KiteOption func(*KiteOptions)

func WithBrand(brand string) KiteOption {
	return func(opt *KiteOptions) {
		opt.Brand = brand
	}
}

func WithColor(color string) KiteOption {
	return func(opt *KiteOptions) {
		opt.Color = color
	}
}

func WithWeight(weight int) KiteOption {
	return func(opt *KiteOptions) {
		opt.Weight = weight
	}
}

func NewKite(options ...KiteOption) *Kite {
	opts := &KiteOptions{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	for _, option := range options {
		option(opts)
	}

	return &Kite{
		Brand:  opts.Brand,
		Color:  opts.Color,
		Weight: opts.Weight,
	}
}
