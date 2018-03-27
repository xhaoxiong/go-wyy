package main

import (
	"go-wyy/v2.0/engin"

	"fmt"
	"encoding/json"
	"go-wyy/models"
	"ccmous/mooc/crawler/engine"
	"io"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	engin.Run(
		engin.Request{
			Url: "",
			ParserComment: engin.NilParser,
			ParserSong: func(reader io.Reader) engin.ParseResult {
				doc, err := goquery.NewDocumentFromReader(reader)

				if err != nil {
					panic(err)
				}

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
					fmt.Printf("歌曲题目:%s\n", title)
					result.Items = append(result.Items, "歌曲信息:", song)
					go func() {

					}()
					//song.UserId = userId


				})
				return result
			},
		})
}

func ParseComment(content []byte) (result engine.ParseResult) {
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
