package main

import (
	"container/list"
	"database/sql"
	"fmt"
	//"strconv"
	"time"

	"github.com/go-gorp/gorp"
	//"github.com/ocraman/concurrent-map"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "ajay"
	DB_PASSWORD = "Iotsc2"
	DB_NAME     = "sctest"
)

var (
	db         *sql.DB
	mapTableId map[int]string
)

//query a db
//func exampleDBQuery() {

//}
func init() {
	mapTableId = make(map[int]string)

	mapTableId[1] = "SELECT * FROM LEVEL1 where pid=$1"
	mapTableId[2] = "SELECT * FROM LEVEL2 where pid=$1"
	mapTableId[3] = "SELECT * FROM LEVEL3 where pid=$1"
}

type Post struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`               // Column size set to 50
	Body    string `db:"article_body,size:1024"` // Set both column name and size
}

func main() {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)
	checkErr(err)
	defer db.Close()

	dbm := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbm.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
	dbm.CreateTables()

	//_, err = db.Exec("INSERT INTO LEVEL1 (ID, NAME) VALUES($1, $2)", 4, "MASOOR")
	//res, err := db.Query("SELECT * FROM level1;")
	//checkErr(err)
	getAllLights(1, 1, db)
}

const bottom = 3

func getAllLights(level, id int, db *sql.DB) {
	plights := list.New()
	plights.PushBack(id)
	//res := make([]int)
	for level < bottom {
		level++
		tlights := list.New()
		for cl := plights.Front(); cl != nil; cl = cl.Next() {
			qstmt := mapTableId[level]
			res, err := db.Query(qstmt, cl.Value)
			checkErr(err)
			for res.Next() {
				var name string
				var id, pid int
				err = res.Scan(&id, &name, &pid)
				checkErr(err)
				tlights.PushBack(id)
				fmt.Println("name = ", name, "id = ", id, "pid = ", pid)
			}
		}
		plights = tlights
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
