package golf

const (
	DefaultBrand  = "DefaultGolfBrand"
	DefaultColor  = "DefaultGolfColor"
	DefaultWeight = 3
)

type GolfClub struct {
	Brand, Color string
	Weight       int
}

type GolfClubBuilder struct {
	club *GolfClub
}

func NewGolfClubBuilder() *GolfClubBuilder {
	return &GolfClubBuilder{
		&GolfClub{
			Brand:  DefaultBrand,
			Color:  DefaultColor,
			Weight: DefaultWeight,
		}}
}

func (gb *GolfClubBuilder) Build() *GolfClub { return gb.club }

func (gb *GolfClubBuilder) Brand(brand string) *GolfClubBuilder {
	gb.club.Brand = brand
	return gb
}

func (gb *GolfClubBuilder) Type(clubType string) *GolfClubBuilder {
	gb.club.Color = clubType
	return gb
}

func (gb *GolfClubBuilder) Length(length int) *GolfClubBuilder {
	gb.club.Weight = length
	return gb
}
