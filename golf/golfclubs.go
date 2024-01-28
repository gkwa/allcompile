package golf

type GolfClubs []*GolfClub

func NewGolfClubs() GolfClubs {
	return make(GolfClubs, 0)
}

func (gc GolfClubs) AddClub(club *GolfClub) GolfClubs {
	return append(gc, club)
}
