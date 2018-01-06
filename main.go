package main

import (

	"go-wyy/service/songs"
	"fmt"
)

func main() {
		data,err:=songs.GetDownloadUrl("[497034035]","320000")
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(data.Data[0].Url)


		//songs.Songs()
		//comment1,err:=comment.GetAllComment("404465600")
		////comment,err:=comment.GetComments("404465600",0,20)
		//if err!=nil{
		//	panic(err)
		//}
		//fmt.Println(comment1)
	}
