package parse

import (
	"fmt"
	"encoding/json"
	"go-wyy/models"
	"ccmous/mooc/crawler/engine"
)

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
