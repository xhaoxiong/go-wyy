package engin

import (
	"log"
	"go-wyy/v2.0/fetcher"
)

//type SimpleEngine struct {}

func Run(seeds ...Request) {
	//result, err := workerComment(seeds[0])
	parseResult, err := workerSong(seeds[0])

	if err != nil {
		panic(err)
	}
	for _,item:=range parseResult.Items {
		log.Printf("Get songItem %v", item)
	}


	//for _, item := range result.Items {
	//	log.Printf("Get commentItem %v\n", item)
	//}
}



func workerSong(r Request) (ParseResult, error) {
	reader, err := fetcher.GetSongs("462312279")
	if err != nil {
		log.Printf("Fetcher:error fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserSong(reader), nil
}
