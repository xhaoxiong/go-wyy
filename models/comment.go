package models

type Comments struct {
	IsMusician  bool          `json:"is_musician"`
	UserId      int64         `json:"user_id"`
	TopComments []string      `json:"top_comments"`
	MoreHot     bool          `json:"more_hot"`
	HotComments []*HotComment `json:"hot_comments"`
}

type HotComment struct {
	User       *User    `json:"user"`
	BeReplied  []string `json:"be_replied"`
	LikedCount int64    `json:"liked_count"`
	Liked      bool     `json:"liked"`
	CommentId  int64    `json:"comment_id"`
	Time       int64    `json:"time"`
	Content    string   `json:"content"`
}

type User struct {
	LocationInfo *LocationInfo `json:"-"`
	UserType     int           `json:"user_type"`
	ExpertTags   *ExpertTag    `json:"-"`
	UserId       int64         `json:"user_id"`
	NickName     string        `json:"nick_name"`
	Experts      *Expert       `json:"-"`
	AuthStatus   int           `json:"auth_status"`
	RemarkName   *RemarkName   `json:"-"`
	AvatarUrl    string        `json:"avatar_url"`
	VipType      int           `json:"vip_type"`
}

type LocationInfo struct {
}

type ExpertTag struct {
}

type Expert struct {
}

type RemarkName struct {
}
