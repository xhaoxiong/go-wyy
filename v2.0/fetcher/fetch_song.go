package fetcher

import (
	"net/http"
	"io"
)

func GetSongs(userId string) (io.Reader,error){
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
	return res.Body,err
}
