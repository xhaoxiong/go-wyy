package parse

import (
	"github.com/PuerkitoBio/goquery"
	"sync"
	"go-wyy/models"
	"io"
	"go-wyy/v2.0/engin"
	"go-wyy/v2.0/fetcher"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"fmt"
)

func ParseSong(reader io.Reader) engin.ParseResult {
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	result := engin.ParseResult{}
	doc.Find("ul[class=f-hide] a").Each(func(i int, selection *goquery.Selection) {
		/*开启协程插入数据库，并且开启协程请求每首歌的评论*/
		songIdUrl, _ := selection.Attr("href")
		title := selection.Text()
		var song models.Song
		//歌曲id
		songId := songIdUrl[9:len(songIdUrl)]
		song.SongId = songId

		///song?id=歌曲id
		song.SongUrlId = songIdUrl

		//歌曲标题
		song.Title = title
		//fmt.Printf("歌曲题目:%s\n", title)
		result.Items = append(result.Items, "歌曲信息:", song)
		offset := 0
		songComment := make(chan []byte, 100)
		go func(offset int) {
			for {
				fetcher.GetComments(songId, offset, offset+40, songComment, wg)
				offset += 40
				wg.Add(1)
			}

		}(offset)

		go ReceiveComment(songComment, wg)
	})

	return result
}

func ReceiveComment(songComment chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if bytes, ok := <-songComment; ok {
			fmt.Println(bytes)
			//fmt.Println(string(bytes))
			save(bytes)
		} else {
			close(songComment)
			break
		}

	}

}

func save(item []byte) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)

	if err != nil {
		panic(err)
	}

		response, err := client.Index().
			Index("wyy").
			Type("comment").
			BodyJson(string(item)).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Println(response)


}
