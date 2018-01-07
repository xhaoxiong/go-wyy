package models

type Song struct {
	Id          int64
	SongUrlId   string
	SongId      string
	Title       string
	DownloadUrl string
	UserId      string
}

type PlayList struct {
	Id     int64
	UserId string
	Songs  []Song
}
