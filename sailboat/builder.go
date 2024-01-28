// sailboat/builder.go
package sailboat

type Sailboat struct {
	Brand  string
	Color  string
	Length int
}

// SailboatBuilder is a builder for creating Sailboat instances.
type SailboatBuilder struct {
	boat *Sailboat
}

// NewSailboatBuilder creates a new SailboatBuilder with default values.
func NewSailboatBuilder() *SailboatBuilder {
	return &SailboatBuilder{boat: &Sailboat{
		Brand:  "Unknown",
		Color:  "Unspecified",
		Length: 999,
	}}
}

// Brand sets the brand property of the sailboat.
func (sb *SailboatBuilder) Brand(brand string) *SailboatBuilder {
	sb.boat.Brand = brand
	return sb
}

// Color sets the color property of the sailboat.
func (sb *SailboatBuilder) Color(color string) *SailboatBuilder {
	sb.boat.Color = color
	return sb
}

// Length sets the length property of the sailboat.
func (sb *SailboatBuilder) Length(length int) *SailboatBuilder {
	sb.boat.Length = length
	return sb
}

// Build creates the final Sailboat instance.
func (sb *SailboatBuilder) Build() *Sailboat {
	return sb.boat
}
