package golf

type GolfClub struct {
	Brand, Type string
	Length      int
}

type GolfClubBuilder struct{ club *GolfClub }

func NewGolfClubBuilder() *GolfClubBuilder {
	return &GolfClubBuilder{&GolfClub{"DefaultBrand", "DefaultType", 999}}
}

func (gb *GolfClubBuilder) Build() *GolfClub { return gb.club }

func (gb *GolfClubBuilder) Brand(brand string) *GolfClubBuilder {
	gb.club.Brand = brand
	return gb
}

func (gb *GolfClubBuilder) Type(clubType string) *GolfClubBuilder {
	gb.club.Type = clubType
	return gb
}

func (gb *GolfClubBuilder) Length(length int) *GolfClubBuilder {
	gb.club.Length = length
	return gb
}
