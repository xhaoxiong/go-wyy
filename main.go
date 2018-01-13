package main

import (
	"go-wyy/models"
	"os"
	"starvote/task"
)


func main() {
	initArgs()

	/**
		歌单所有评论思路：进入个人歌单，获取歌曲id，爬取评论


	*/

	//评论测试
	//data,err:=comment.GetComments("460628744", 0, 40)
	//
	//if err!=nil{
	//	panic(err)
	//}
	//fmt.Println("isMusician",data.IsMusician)
	//fmt.Println("Total",data.Total)
	//fmt.Println("HotComment",data.HotComments)
	//fmt.Println("Code",data.Code)
	//fmt.Println("More Hot",data.MoreHot)
	//
	//
	models.Connect()


	//songs.Songs("462312279")
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
