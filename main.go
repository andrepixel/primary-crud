package main

import (
	_ "github.com/mattn/go-sqlite3"

	"github.com/flukebr/primary-crud/database"
	"github.com/flukebr/primary-crud/entity"
)

const (
	FILE_DB   = "./myData.db"
	DRIVER_DB = "sqlite3"
)

func main() {
	defer func() {
        if r := recover(); r != nil {
            println("Recovered. Error:\n", r)
        }
    }()

	game := entity.Game{}

	db := database.OpenConnectionWithDatabase(DRIVER_DB, FILE_DB)

	database.CreateTable(db, `
		create table if not exists 'Games'(
			'id' varchar(16) primary key,
			'name' varchar(60) not null,
			'dateOfExist' integer not null,
			'company' varchar(60) not null,
			'releaseYear' integer not null
		);
	`)

	newGame := game.InsertGame("Naruto", "Bandai", 2000)

	database.InsertTable(db, "insert into 'Games' values (?,?,?,?,?)", newGame.GetIdString(), newGame.GetName(), newGame.GetCompany(), newGame.GetReleaseYear(), game.GetDateOfExist())
	
	newGame2 := game.InsertGame("TLOU", "Sony", 2010)

	database.InsertTable(db, "insert into 'Games' values (?,?,?,?,?)", newGame2.GetIdString(), newGame2.GetName(), newGame2.GetCompany(), newGame2.GetReleaseYear(), game.GetDateOfExist())

	newGame3 := game.InsertGame("Halo", "Microsoft", 2021)
	
	database.InsertTable(db, "insert into 'Games' values (?,?,?,?,?)", newGame3.GetIdString(), newGame3.GetName(), newGame3.GetCompany(), newGame3.GetReleaseYear(), game.GetDateOfExist())

	database.UpdateData(db, "update 'Games' set name = ? where name = ?;", "Halo Infinite", "Halo")

	database.RemoveData(db, "delete from 'Games' where name = ?;", "aa")

	result := database.GetAllDataInTable(db, "select * from 'Games';")

	database.PrintRowsInConsole(result, newGame)

	database.CloseDB(db)
	database.RemoveDB(FILE_DB)
}
