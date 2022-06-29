package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/flukebr/primary-crud/entity"
	_ "github.com/mattn/go-sqlite3"
)

func OpenConnectionWithDatabase(driverDB, fileDB string) *sql.DB {
	if _, err := os.Stat(fileDB); os.IsNotExist(err) {
		db, err := sql.Open(driverDB, fileDB)

		if err != nil {
			fmt.Println(err)
		}

		return db
	} else {
		RemoveDB(fileDB)

		db, err := sql.Open(driverDB, fileDB)

		if err != nil {
			fmt.Println(err)
		}

		return db
	}
}

func CreateTable(db *sql.DB, query string) {
	stmt, err := db.Prepare(query)

	if err != nil {
		fmt.Println(err)
	}

	stmt.Exec()
}

func InsertTable(db *sql.DB, query string, args ...interface{}) {
	result, err := db.Prepare(query)

	if err != nil {
		fmt.Println(err)
		recover()
	}

	result.Exec(args...)
}

func GetAllDataInTable(db *sql.DB, query string) *sql.Rows {
	result, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func UpdateData(db *sql.DB, query string, args ...interface{}) {
	result, err := db.Prepare(query)

	if err != nil {
		fmt.Println(err)
	}

	result.Exec(args...)
}

func RemoveData(db *sql.DB, query string, args ...interface{}) {
	result, err := db.Prepare(query)

	if err != nil {
		fmt.Println(err)
	}

	result.Exec(args...)
}

func PrintRowsInConsole(data *sql.Rows, game *entity.Game) {
	fmt.Println("---------------------------")

	for data.Next() {
		data.Scan(&game.Id, &game.Name, &game.Company, &game.ReleaseYear, &game.DateOfExist)

		fmt.Printf("'id': %s\n'name': %s\n'dateofExist': %d years\n'company': %s\n'releaseYear': %d\n\n", game.GetIdString(), game.GetName(), game.GetDateOfExist(), game.GetCompany(), game.GetReleaseYear())
	}

}

func CloseDB(db *sql.DB) {
	db.Close()
}

func RemoveDB(pathFile string) {
	os.Remove(pathFile)
}
