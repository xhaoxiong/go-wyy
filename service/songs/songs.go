package songs

import (
	"net/http"
	"github.com/github.com/PuerkitoBio/goquery"
	"fmt"
	"go-wyy/models"
	"go-wyy/service/comment"
)

func getSongList() {

}

func Songs() {
	userId := "317113395"
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
	var playList models.PlayList
	var songs []models.Song
	var song models.Song
	doc.Find("ul[class=f-hide] a").Each(func(i int, selection *goquery.Selection) {
		/*开启协程插入数据库，并且开启协程请求每首歌的评论*/
		songIdUrl, _ := selection.Attr("href")
		title := selection.Text()
		songId:=songIdUrl[9:len(songIdUrl)]
		song.SongId=songId
		song.SongUrlId=songIdUrl
		song.Title=title
		songs=append(songs,song)
		go comment.GetAllComment(songId)
	})

	for _,v:=range songs{
		fmt.Println(v.Title)
		fmt.Println(v.SongUrlId)
		fmt.Println(v.SongId)
	}
	playList.UserId=userId
	playList.Songs=songs
}
