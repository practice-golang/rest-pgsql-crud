package dbbooks

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

// Book - A book info
type Book struct {
	ID     int
	Title  string
	Author string
}

func dbConn() (db *sql.DB) {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "root"
	dbPassword := ""
	dbName := "postgres"

	dbinfo := fmt.Sprintf(
		"host='%s' port='%s' user='%s' password='%s' dbname='%s' sslmode='disable'",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// CreateTable - Book Table creation
func CreateTable(table string) (err error) {
	db := dbConn()

	que := `
	CREATE TABLE IF NOT EXISTS "` + table + `"
	(
		"_id" serial NOT NULL,
		"title" character varying(255) NOT NULL,
		"author" character varying(255) NOT NULL,
		-- "created" date,
		-- created_at timestamp with time zone DEFAULT current_timestamp,
		CONSTRAINT userinfo_pkey PRIMARY KEY ("_id")
	)
	-- ) WITH (OIDS=FALSE); // Not work with CockroachDB`

	_, err = db.Exec(que)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// SelectData : cRud
func SelectData(id int, table string) (result []Book) {
	db := dbConn()
	defer db.Close()

	sql := `SELECT * FROM "` + table + `"`
	var where string

	if id > 0 {
		where = ` WHERE "_id"=` + strconv.Itoa(id)
	} else {
		where = ``
	}

	order := ` order by "_id" desc`
	sql = sql + where + order

	rows, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	book := Book{}
	result = []Book{}

	for rows.Next() {
		var _id int
		var title, author string

		err = rows.Scan(&_id, &title, &author)
		if err != nil {
			panic(err.Error())
		}

		book.ID = _id
		book.Title = title
		book.Author = author
		result = append(result, book)
	}

	return
}

// InsertData : Crud
func InsertData(book *Book, table string) error {
	db := dbConn()
	defer db.Close()

	var _id int

	err := db.QueryRow(
		`INSERT INTO "`+table+`"("title","author") VALUES($1,$2) RETURNING _id`,
		book.Title, book.Author).Scan(&_id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// UpdateData : crUd
func UpdateData(book *Book, table string) {
	db := dbConn()
	defer db.Close()

	_, err := db.Exec(
		`UPDATE "`+table+`" SET "title"=$1,"author"=$2 WHERE "_id"=$3 RETURNING "_id"`,
		book.Title, book.Author, book.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteData : cruD
func DeleteData(_id int, table string) {
	db := dbConn()
	defer db.Close()

	_, err := db.Exec(`DELETE FROM "`+table+`" WHERE "_id"=$1`, _id)
	if err != nil {
		log.Fatal(err)
	}
}
