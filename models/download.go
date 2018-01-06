package models

type DownloadData struct {
	Data []*Data `json:"data"`
	Code int     `json:"code"`
}

type Data struct {
	Id        int64   `json:"id"`
	Url       string  `json:"url"`
	Br        int64   `json:"br"`
	Md5       string  `json:"md_5"`
	Code      int     `json:"code"`
	Expi      int     `json:"expi"`
	Type      string  `json:"type"`
	Gain      float64 `json:"gain"`
	Fee       int     `json:"fee"`
	Uf        *Uf     `json:"-"`
	Payed     int     `json:"payed"`
	Flag      int     `json:"flag"`
	CanExtend bool    `json:"can_extend"`
}

type Uf struct {
}
