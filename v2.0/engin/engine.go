package engin

import (
	"log"
	"go-wyy/v2.0/fetcher"
)

//type SimpleEngine struct {}


func Run(seeds ...Request){
	result, err := worker(seeds[0])
	if err!=nil{
		panic(err)
	}
	for _,item:=range result.Items{
		log.Printf("Get item %v", item)
	}
}


func worker(r Request) (ParseResult, error) {
	body, err := fetcher.GetComments("460628744", 0, 40)
	if err != nil {
		log.Printf("Fetcher:error fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser(body), nil
}
