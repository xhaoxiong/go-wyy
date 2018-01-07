package models

type Commentt struct {
	Id         int64
	IsMusician bool  `json:"isMusician"`
	UserId     int32 `json:"userId"`
	//TopComments []string       `json:"topComments";gorm:"-"`
	MoreHot     bool           `json:"moreHot"`
	HotComments []*HotComments `json:"hotComments"`
	Code        int            `json:"code"`
	Comments    []*Comments    `json:"comments"`
	Total       int64          `json:"total"`
	More        bool           `json:"more"`
	SongId      string
}

type Comments struct {
	Id                 int64
	User               *User        `json:"user"`
	BeReplied          []*BeReplied `json:"-"`
	Time               int64        `json:"time"`
	LikedCount         int          `json:"likedCount"`
	Liked              bool         `json:"liked"`
	CommentId          int64        `json:"commentId"`
	Content            string       `json:"content";gorm:"type:longtext"`
	IsRemoveHotComment bool         `json:"isRemoveHotComment"`
	Commentt           *Commentt
	CommenttID         int64
}

type HotComments struct {
	Id         int64
	User       *User        `json:"user"`
	BeReplied  []*BeReplied `json:"-"`
	Time       int64        `json:"time"`
	LikedCount int          `json:"likedCount"`
	Liked      bool         `json:"liked"`
	CommentId  int64        `json:"commentId"`
	Content    string       `json:"content";gorm:"type:longtext"`
	Commentt   *Commentt
	CommenttID int64
}

type User struct {
	Id           int64
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
	Comments     *Comments
	CommentsID   int64
}

type BeReplied struct {
	Id            int64
	User          *User  `json:"-"`
	UserID        int64
	Content       string `json:"content"`
	Status        int    `json:"status"`
	CommentsID    int64
	HotCommentsID int64
}

type LocationInfo struct {
}

type ExpertTag struct {
}

type Expert struct {
}

type RemarkName struct {
}
