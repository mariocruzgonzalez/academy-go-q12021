package model

type ChuckJoke struct {
	ID			int 		`gorm:"primary_key" json:"id"`
	Joke 		string 		`json:"joke"`
}
