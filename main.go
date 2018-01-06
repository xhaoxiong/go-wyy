package main

import (

	"go-wyy/service/songs"
	"go-wyy/models"
	"os"
	"starvote/task"
	"time"
)

func main() {
		//data,err:=songs.GetDownloadUrl("[497034035]","320000")
		//if err!=nil{
		//	fmt.Println(err)
		//}
		//fmt.Println(data.Data[0].Url)
		initArgs()
		models.Connect()

		songs.Songs("462312279")
	time.Sleep(10000000*time.Second)
		//comment1,err:=comment.GetAllComment("404465600")
		////comment,err:=comment.GetComments("404465600",0,20)
		//if err!=nil{
		//	panic(err)
		//}
		//fmt.Println(comment1)
	}


func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.SyncDB()
			os.Exit(0)
		}
		if v == "-admin" {
			models.Connect()
			models.AddAdmin()
			os.Exit(0)
		}
		if v == "-clean" {
			models.Connect()
			task.CleanRepeat()
			os.Exit(0)
		}
	}
}