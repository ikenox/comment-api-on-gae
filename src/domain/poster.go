package domain

type Poster struct {
	Entity
	posterId int64
	profile  PosterProfile
}

type PosterProfile struct {
	ValueObject
	name string
}
