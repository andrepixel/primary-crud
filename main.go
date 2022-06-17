package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/flukebr/primary-crud/entity"
)

const (
	FILE_DB   = "./myData.db"
	DRIVER_DB = "sqlite3"
)

func main() {
	game := entity.Game{}

	db := openConnectionWithDatabase(DRIVER_DB, FILE_DB)

	createTable(db, `
		create table if not exists 'Games'(
			'id' integer primary key autoincrement,
			'name' varchar(60) not null,
			'dateOfExist' varchar(60) not null,
			'company' varchar(60) not null,
			'releaseYear' varchar(60) not null
		);
	`)

	insertTable(db, "insert into 'Games' values (1,'Naruto','2000','Bandai','2022');")

	result := getAllTable(db, "select * from Games;")

	printDataInConsole(result, &game)

	closeDB(db)
	removeDB(FILE_DB)
}

func openConnectionWithDatabase(driverDB, fileDB string) *sql.DB {
	if _, err := os.Stat(fileDB); os.IsNotExist(err) {
		db, err := sql.Open(driverDB, fileDB)

		if err != nil {
			fmt.Println(err)
		}

		return db
	} else {
		removeDB(fileDB)
		return nil
	}
}

func createTable(db *sql.DB, query string) {
	stmt, err := db.Prepare(query)

	if err != nil {
		fmt.Println(err)
	}

	stmt.Exec()
}

func insertTable(db *sql.DB, query string) {
	result, err := db.Prepare(query)

	if err != nil {
		fmt.Println(err)
	}

	result.Exec()
}

func getAllTable(db *sql.DB, query string) *sql.Rows {
	result, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func printDataInConsole(data *sql.Rows, game *entity.Game) {
	for data.Next() {
		data.Scan(&game.Id, &game.Name, &game.DateOfExist, &game.Company, &game.ReleaseYear)
		fmt.Printf("'id': %d\n'name': %s\n'dateofExist': %s\n'company': %s\n'releaseYear': %s\n", game.Id, game.Name, game.DateOfExist, game.Company, game.ReleaseYear)
	}
}

func closeDB(db *sql.DB) {
	db.Close()
}

func removeDB(pathFile string) {
	os.Remove(pathFile)

}
