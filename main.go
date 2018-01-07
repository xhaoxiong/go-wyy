package main

import (
	"go-wyy/models"
	"os"
	"starvote/task"
	"go-wyy/service/songs"
)

func main() {
	initArgs()
	db_name := "root"
	db_pass := "971129XLZ"
	models.Connect(db_name, db_pass)
	songs.Songs("462312279")
}

func initArgs() {
	args := os.Args
	db_name := "root"
	db_pass := "971129XLZ"
	for _, v := range args {
		if v == "-syncdb" {
			models.SyncDB(db_name,db_pass)
			os.Exit(0)
		}
		if v == "-admin" {
			models.Connect(db_name, db_pass)
			models.AddAdmin()
			os.Exit(0)
		}
		if v == "-clean" {
			models.Connect(db_name, db_pass)
			task.CleanRepeat()
			os.Exit(0)
		}
	}
}
