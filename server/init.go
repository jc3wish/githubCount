package server

import (
	"database/sql"
	"fmt"
	"github.com/jc3wish/githubCount/config"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"sync"
)

var db *sql.DB

var l sync.RWMutex
func GetDbConn() *sql.DB {
	l.Lock()
	return db
}

func BackDbConn(db *sql.DB)  {
	l.Unlock()
}

var sqlArr = []string{
	//"CREATE DATABASE IF NOT EXISTS myGithubCount",
	`CREATE TABLE IF NOT EXISTS  x_project(
		   id INTEGER PRIMARY KEY AUTOINCREMENT,
		   name           TEXT    NOT NULL,
		   add_time       INT     NOT NULL
		)`,
	`CREATE TABLE  IF NOT EXISTS  x_project_star(
		   id INTEGER PRIMARY KEY AUTOINCREMENT,
		   project_id           INT    NOT NULL,
		   add_time       INT     NOT NULL,
			subscribers_count INT NOT NULL,
			stargazers_count INT NOT NULL,
			forks_count INT NOT NULL
		)`,
		`CREATE UNIQUE INDEX IF NOT EXISTS  x_project_name on x_project (name)`,
	`CREATE INDEX IF NOT EXISTS  x_project_star_index on x_project_star (project_id)`,
}

func DoInit()  {
	os.Mkdir(config.DataDir,0700)
	var err error
	db, err = sql.Open("sqlite3", config.DataDir+"/sqlite3.db")
	if err != nil {
		fmt.Println("non sqlite3.db", err)
		log.Fatal(err)
	}
	for _,sql := range sqlArr{
		_,err = db.Exec(sql)
		if err != nil {
			log.Fatal(err)
		}
	}

	StartCron()
}