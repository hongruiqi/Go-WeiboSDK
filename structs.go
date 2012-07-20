package weibo

type APIError struct {
	Request          string `json:"request"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

func (e *APIError) Error() string {
	return e.ErrorDescription
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	RemindIn    string `json:"remind_in"`
	Uid         string `json:"uid"`
}

type FriendsTimeline struct {
	Statuses       []Status      `json:"statuses"`
	Advertises     []interface{} `json:"advertises"`
	HasVisible     bool          `json:"has_visible"`
	PreviousCursor int64         `json:"previous_cursor"`
	NextCursor     int64         `json:"next_cursor"`
	TotalNumber    int64         `json:"total_number"`
}

type HomeTimeline FriendsTimeline

type FriendsTimelineIds struct {
	Statuses       []string      `json:"statuses"`
	Advertises     []interface{} `json:"advertises"`
	HasVisible     bool          `json:"has_visible"`
	PreviousCursor int64         `json:"previous_cursor"`
	NextCursor     int64         `json:"next_cursor"`
	TotalNumber    int64         `json:"total_number"`
}

type PublicTimeline struct {
	Statuses       []Status `json:"statuses"`
	HasVisible     bool     `json:"has_visible"`
	PreviousCursor int64    `json:"previous_cursor"`
	NextCursor     int64    `json:"next_cursor"`
	TotalNumber    int64    `json:"total_number"`
}

type UserTimeline struct {
	Statuses       []Status      `json:"statuses"`
	Marks          []interface{} `json:"marks"`
	HasVisible     bool          `json:"has_visible"`
	PreviousCursor int64         `json:"previous_cursor"`
	NextCursor     int64         `json:"next_cursor"`
	TotalNumber    int64         `json:"total_number"`
}

type UserTimelineIds struct {
	Statuses       []string      `json:"statuses"`
	Marks          []interface{} `json:"marks"`
	HasVisible     bool          `json:"has_visible"`
	PreviousCursor int64         `json:"previous_cursor"`
	NextCursor     int64         `json:"next_cursor"`
	TotalNumber    int64         `json:"total_number"`
}

type RepostTimeline struct {
	Reposts        []Status `json:"reposts"`
	HasVisible     bool     `json:"has_visible"`
	PreviousCursor int64    `json:"previous_cursor"`
	NextCursor     int64    `json:"next_cursor"`
	TotalNumber    int64    `json:"total_number"`
}

type RepostTimelineIds struct {
	Statuses       []string `json:"statuses"`
	HasVisible     bool     `json:"has_visible"`
	PreviousCursor int64    `json:"previous_cursor"`
	NextCursor     int64    `json:"next_cursor"`
	TotalNumber    int64    `json:"total_number"`
}

type RepostByMe RepostTimeline

type StatusMentions PublicTimeline

type StatusMentionsIds struct {
	Statuses       []string `json:"statuses"`
	HasVisible     bool     `json:"has_visible"`
	PreviousCursor int64    `json:"previous_cursor"`
	NextCursor     int64    `json:"next_cursor"`
	TotalNumber    int64    `json:"total_number"`
}

type BilateralTimeline PublicTimeline

type Querymid_One struct {
	Mid string `json:"mid"`
}

type Querymid_Batch []map[string]string

type HotRepostDaily []Status
type HotRepostWeekly []Status
type HotCommentsDaily []Status
type HotCommentsWeekly []Status

type Count []struct {
	Id       int64 `json:"id"`
	Comments int   `json:"comments"`
	Reposts  int   `json:"reposts"`
}

type Visible struct {
	Type   int `json:"type"`
	ListId int `json:"list_id"`
}

type Status struct {
	CreatedAt           string        `json:"created_at"`
	Id                  int64         `json:"id"`
	Mid                 string        `json:"mid"`
	IdStr               string        `json:"idstr"`
	Text                string        `json:"text"`
	Source              string        `json:"source"`
	Favorited           bool          `json:"favorited"`
	Truncated           bool          `json:truncated"`
	InReplyToStatusId   string        `json:"in_reply_to_status_id"`
	InReplyToUserId     string        `json:"in_reply_to_user_id"`
	InReplyToScreenName string        `json:"in_reply_to_screen_name"`
	ThumbnailPic        string        `json:"thumbnail_pic"`
	BmiddlePic          string        `json:"bmiddle_pic"`
	OriginalPic         string        `json:"original_pic"`
	Geo                 *Geo          `json:"geo"`
	User                *User         `json:"user"`
	RetweetedStatus     *Status       `json:"retweeted_status"`
	RepostsCount        int           `json:"reposts_count"`
	CommentsCount       int           `json:"comments_count"`
	Annotations         []interface{} `json:"annotations"`
	Mlevel              int           `json:"mlevel"`
	Visible             Visible       `json:"visible"`
}

type User struct {
	Id               int64  `json:"id"`
	IdStr            string `json:"idstr"`
	ScreenName       string `json:"screen_name"`
	Name             string `json:"name"`
	Province         string `json:"province"`
	City             string `json:"city"`
	Location         string `json:"location"`
	Description      string `json:"description"`
	Url              string `json:"url"`
	ProfileImageUrl  string `json:"profile_image_url"`
	ProfileUrl       string `json:"profile_url"`
	Domain           string `json:"domain"`
	Weihao           string `json:"weihao"`
	Gender           string `json:"gender"`
	FollowersCount   int    `json:"followers_count"`
	StatusesCount    int    `json:"statuses_count"`
	FavouritesCount  int    `json:"favourites_count"`
	CreatedAt        string `json:"created_at"`
	Following        bool   `json:"following"`
	AllowAllActMsg   bool   `json:"allow_all_act_msg"`
	GeoEnabled       bool   `json:"geo_enabled"`
	Verified         bool   `json:"verified"`
	VerifiedType     int    `json:"verified_type"`
	Remark           string `json:"remark"`
	AllowAllComment  bool   `json:"allow_all_comment"`
	AvatarLarge      string `json:"avatar_large"`
	VerifiedReason   string `json:"verified_reason"`
	FollowMe         bool   `json:"follow_me"`
	OnlineStatus     int    `json:"online_status"`
	BiFollowersCount int    `json:"bi_followers_count"`
	Lang             string `jsson:"lang"`
}

type Geo struct {
	Type        string     `json:"Point"`
	Coordinates [2]float64 `json:"coordinates"`
}

type Emotion struct {
	Phrase   string `json:"phrase"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Hot      bool   `json:"hot"`
	Common   bool   `json:"common"`
	Category string `json:"category"`
	Icon     string `json:"icon"`
	Value    string `json:"value"`
	Picid    string `json:"picid"`
}

type Emotions []Emotion
