package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Game struct {
	Id          uuid.UUID
	Name        string
	DateOfExist int
	Company     string
	ReleaseYear int
}

func (game *Game) InsertGame(name, company string, releaseYear int) *Game {
	game.Id = uuid.NewV4()
	game.SetName(name)
	game.SetCompany(company)
	game.SetReleaseYear(releaseYear)
	game.calculateDateOfExist(time.Now().Year())

	return game
}

func (game *Game) calculateDateOfExist(date int) {
	game.DateOfExist = date - game.ReleaseYear
}

func (game *Game) GetIdString() string {
	return game.Id.String()
}

func (game *Game) GetName() string {
	return game.Name
}

func (game *Game) GetDateOfExist() int {
	return game.DateOfExist
}

func (game *Game) GetCompany() string {
	return game.Company
}

func (game *Game) GetReleaseYear() int {
	return game.ReleaseYear
}

func (game *Game) SetName(Name string) *Game {
	game.Name = Name
	return game
}

func (game *Game) SetDateOfExist(dateOfExist int) *Game {
	game.DateOfExist = dateOfExist
	return game
}

func (game *Game) SetCompany(company string) *Game {
	game.Company = company
	return game
}

func (game *Game) SetReleaseYear(releaseYear int) *Game {
	game.ReleaseYear = releaseYear
	return game
}
