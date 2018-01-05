package models

type Comment struct {
	IsMusician  bool           `json:"isMusician"`
	UserId      int32          `json:"userId"`
	TopComments []string       `json:"topComments"`
	MoreHot     bool           `json:"moreHot"`
	HotComments []*HotComments `json:"hotComments"`
	Code        int            `json:"code"`
	Comments    []*Comments    `json:"comments"`
	Total       int64          `json:"total"`
	More        bool           `json:"more"`
}

type Comments struct {
	User               *User        `json:"user"`
	BeReplied          []*BeReplied `json:"beReplied"`
	Time               int64        `json:"time"`
	LikedCount         int          `json:"likedCount"`
	Liked              bool         `json:"liked"`
	CommentId          int64        `json:"commentId"`
	Content            string       `json:"content"`
	IsRemoveHotComment bool         `json:"isRemoveHotComment"`
}

type HotComments struct {
	User       *User        `json:"user"`
	BeReplied  []*BeReplied `json:"beReplied"`
	Time       int64        `json:"time"`
	LikedCount int          `json:"likedCount"`
	Liked      bool         `json:"liked"`
	CommentId  int64        `json:"commentId"`
	Content    string       `json:"content"`
}

type User struct {
	LocationInfo *LocationInfo `json:"-"`
	UserType     int           `json:"userType"`
	ExpertTags   *ExpertTag    `json:"-"`
	UserId       int64         `json:"userId"`
	NickName     string        `json:"nickName"`
	Experts      *Expert       `json:"-"`
	AuthStatus   int           `json:"authStatus"`
	RemarkName   *RemarkName   `json:"-"`
	AvatarUrl    string        `json:"avatarUrl"`
	VipType      int           `json:"vipType"`
}

type BeReplied struct {
	User    *User  `json:"user"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}

type LocationInfo struct {
}

type ExpertTag struct {
}

type Expert struct {
}

type RemarkName struct {
}
