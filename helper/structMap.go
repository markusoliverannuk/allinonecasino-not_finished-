package helper

type User struct {
	Id           string
	Name         string
	Email        string
	PassWordHash string
	Role         int // 0-visitor 1 -user 2-admin
	Balance      int
}

type HomePage struct {
	CurrentUser     User
	Data            []HomePageData
	NotificationsPl []Notifications
}

type Notifications struct {
	User_id      string
	Post_id      string
	ReactionType string
	IsRead       bool
	Post_name    string
}

type HomePageData struct {
	Username           string
	Post_id            string
	Post_userid        string
	Post_content       string
	Post_date          string
	Post_title         string
	Post_like_count    int
	Post_dislike_count int
	Like_active        bool
	Dislike_active     bool
}
