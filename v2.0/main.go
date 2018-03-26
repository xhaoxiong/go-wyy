package main

import (
	"go-wyy/v2.0/engin"

	"fmt"
	"encoding/json"
	"go-wyy/models"
	"ccmous/mooc/crawler/engine"
)

func main()  {

	engin.Run(
		engin.Request{
			Url:"",
			Parser: func(bytes []byte)(result engin.ParseResult) {

				var comment *models.Commentt
				err := json.Unmarshal(bytes, &comment)
				if err != nil {
					fmt.Println(err)

				}
				result=engin.ParseResult{Items:[]interface{}{comment}}
				return
			},
			})
}


func ParseComment(content []byte) (result engine.ParseResult ){
	var comment *models.Commentt
	err := json.Unmarshal(content, &comment)
	if err != nil {
		fmt.Println(err)

	}

	result = engine.ParseResult{
		Items: []interface{}{comment},
	}

	return
}
