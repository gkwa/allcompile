package golf

type GolfClubs []*GolfClub

func (gc GolfClubs) AddClub(club *GolfClub) GolfClubs {
	return append(gc, club)
}
