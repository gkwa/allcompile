package sailboat

const (
	DefaultBrand  = "DefaultSailboatBrand"
	DefaultColor  = "DefaultSailboatColor"
	DefaultWeight = 11
)

type Sailboat struct {
	Brand  string
	Color  string
	Weight int
}

type SailboatBuilder struct {
	boat *Sailboat
}

func NewSailboatBuilder() *SailboatBuilder {
	return &SailboatBuilder{boat: &Sailboat{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}}
}

func (sb *SailboatBuilder) Build() *Sailboat {
	return sb.boat
}

func (sb *SailboatBuilder) Brand(brand string) *SailboatBuilder {
	sb.boat.Brand = brand
	return sb
}

func (sb *SailboatBuilder) Color(color string) *SailboatBuilder {
	sb.boat.Color = color
	return sb
}

func (sb *SailboatBuilder) Length(length int) *SailboatBuilder {
	sb.boat.Weight = length
	return sb
}
