package main

import (
	"go-wyy/service/comment"
	"fmt"
)

func main() {
		commentsJson:=comment.GetComments()
		fmt.Println(commentsJson)
	}
