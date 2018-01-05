package main

import (
	"go-wyy/service/comment"
	"fmt"
)

func main() {
		comment,err:=comment.GetComments("404465600",0,20)
		if err!=nil{
			panic(err)
		}
		fmt.Println(comment.Total)
	}
