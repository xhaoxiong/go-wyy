package songs

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"go-wyy/models"
	"go-wyy/service/comment"
	"sync"
)

/**
	userId 用户Id

*/
func Songs(userId string) {
	req, err := http.NewRequest("GET", "http://music.163.com/playlist?id="+userId, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	req.Header.Set("Referer", "http://music.163.com/")
	req.Header.Set("Host", "music.163.com")

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}

	g := 0
	wg := &sync.WaitGroup{}

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


		song.UserId = userId

		go comment.GetAllComment(songId, wg)

		g++
		wg.Add(1)
	})
	wg.Wait()
}
