package parse

import (
	"fmt"
	"encoding/json"
	"go-wyy/models"
	"go-wyy/v2.0/engin"
)




func ParseComment(content []byte) (result engin.ParseResult){
	var comment *models.Commentt
	err := json.Unmarshal(content, &comment)
	if err != nil {
		fmt.Println(err)

	}

	result = engin.ParseResult{
		Items: []interface{}{comment},
	}

	return
}
