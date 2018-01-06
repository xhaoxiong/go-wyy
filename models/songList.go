package models

type Song struct {
	Id          int64
	SongUrlId   string
	SongId      string
	Title       string
	DownloadUrl string
	PlayListID  int64
}

type PlayList struct {
	Id     int64
	UserId string
	Songs  []Song
}
