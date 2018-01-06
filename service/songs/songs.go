package songs

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"go-wyy/models"
	"go-wyy/service/comment"
	"github.com/astaxie/beego"
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
	var playList models.PlayList
	var songs []models.Song
	var song models.Song

	//medium:=make(chan string,10)
	doc.Find("ul[class=f-hide] a").Each(func(i int, selection *goquery.Selection) {
		/*开启协程插入数据库，并且开启协程请求每首歌的评论*/
		songIdUrl, _ := selection.Attr("href")
		title := selection.Text()
		songId := songIdUrl[9:len(songIdUrl)]
		song.SongId = songId
		song.SongUrlId = songIdUrl
		song.Title = title
		download, err := GetDownloadUrl(songId, "320000")
		if err != nil {
			panic(err)
		}
		if len(download.Data) != 0 {
			song.DownloadUrl = download.Data[0].Url
		}
		songs = append(songs, song)

		//medium<-songId
		//fmt.Println(<-medium)
		go comment.GetAllComment(songId)
	})

	for _, v := range songs {
		song.Id = 0
		song.SongId = v.SongId
		song.DownloadUrl = v.DownloadUrl
		song.Title = v.Title
		song.SongUrlId = v.SongUrlId
		if err := models.DB.Create(&song).Error; err != nil {
			beego.Debug(err)
		}

	}
	playList.UserId = userId
	playList.Songs = songs
	//将歌单插入数据库

}
