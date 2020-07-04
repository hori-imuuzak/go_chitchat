package model

type Post struct {
	Uuid string,
	Body string,
	User User,
	Thread Thread,
}