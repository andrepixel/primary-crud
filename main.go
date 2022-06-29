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
	game := entity.Game{}

	db := database.OpenConnectionWithDatabase(DRIVER_DB, FILE_DB)

	database.CreateTable(db, `
		create table if not exists 'Games'(
			'id' integer primary key autoincrement,
			'name' varchar(60) not null,
			'dateOfExist' varchar(60) not null,
			'company' varchar(60) not null,
			'releaseYear' varchar(60) not null
		);
	`)

	database.InsertTable(db, "insert into 'Games' values (1,'Naruto','2000','Bandai','2022');")
	database.InsertTable(db, "insert into 'Games' values (2,'TLOU','2000','Sony','2010');")
	database.InsertTable(db, "insert into 'Games' values (3,'Halo Infinite','2000','Microsoft','2021');")

	result := database.GetAllDataInTable(db, "select * from Games;")

	database.PrintRowsInConsole(result, &game)

	database.UpdateData(db, "update Games set name = 'Halo' where id = 3;")

	result2 := database.GetAllDataInTable(db, "select * from Games;")

	database.PrintRowsInConsole(result2, &game)

	database.RemoveData(db, "delete from 'Games' where id = 2;")

	result3 := database.GetAllDataInTable(db, "select * from Games;")

	database.PrintRowsInConsole(result3, &game)

	database.CloseDB(db)
	database.RemoveDB(FILE_DB)
}
