package kite

type Kite struct {
	Brand  string
	Color  string
	Length int
}

type KiteOptions struct {
	Brand  string
	Color  string
	Length int
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

func WithLength(length int) KiteOption {
	return func(opt *KiteOptions) {
		opt.Length = length
	}
}

func NewKite(options ...KiteOption) *Kite {
	opts := &KiteOptions{
		Brand:  "DefaultBrand",
		Color:  "DefaultColor",
		Length: 999,
	}

	for _, option := range options {
		option(opts)
	}

	return &Kite{
		Brand:  opts.Brand,
		Color:  opts.Color,
		Length: opts.Length,
	}
}
