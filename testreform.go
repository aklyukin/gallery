package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	reform "gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	_ "github.com/lib/pq"
)

var db *reform.DB

func reforminitDB(config instanceConfig) {

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Pgsql.Host, config.Pgsql.Port,
		config.Pgsql.Username, config.Pgsql.Password, config.Pgsql.Database)
	log.Println(dbinfo)
	conn, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Use new *log.Logger for logging.
	logger := log.New(os.Stderr, "SQL: ", log.Flags())
	db = reform.NewDB(conn, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))

	// err = conn.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("conection to DB success")
	}

	// db.MustExec(schema)
	// user := &User{
	// 	Email:    "test",
	// 	Password: "test2",
	// }

	// if err := db.Save(user); err != nil {
	// 	log.Fatal(err)
	// }

	// CreateUser create user

}

func SaveUser(user *User) {
	myuser3, err := db.FindByPrimaryKeyFrom(UserTable, 3)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("test1")
	log.Println(myuser3.(*User).Email)
	log.Println("test2")
	if err := db.Save(user); err != nil {
		log.Fatal(err)
	}
	log.Println("create user")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
