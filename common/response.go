package common

// 返回响应时需要用到的结构体

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type User struct {
	Id              uint   `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FollowCount     uint   `json:"follow_count,omitempty"`
	FollowerCount   uint   `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	Signature       string `json:"signature,omitempty"`
	TotalFavorited  uint   `json:"total_favorited,omitempty"`
	WorkCount       uint   `json:"work_count,omitempty"`
	FavoriteCount   uint   `json:"favorite_count,omitempty"`
}

type Video struct {
	Id            uint   `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount uint   `json:"favorite_count,omitempty"`
	CommentCount  uint   `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type Comment struct {
	Id         uint   `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type Message struct {
	Id         uint   `json:"id,omitempty"`
	ToUserId   uint   `json:"to_user_id,omitempty"`
	FromUserId uint   `json:"from_user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}
