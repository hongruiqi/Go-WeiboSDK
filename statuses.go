package weibo

import (
	"fmt"
	"mime/multipart"
	"bytes"
)

type statuses struct {
	weibo *Weibo
}

func (self *statuses) PublicTimeline(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/public_timeline.json", access_token, nil, options)
	return self.weibo.get(url, new(PublicTimeline))
}

func (self *statuses) FriendsTimeline(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/friends_timeline.json", access_token, nil, options)
	return self.weibo.get(url, new(FriendsTimeline))
}

func (self *statuses) HomeTimeline(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/home_timeline.json", access_token, nil, options)
	return self.weibo.get(url, new(HomeTimeline))
}

func (self *statuses) FriendsTimelineIds(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/friends_timeline/ids.json", access_token, nil, options)
	return self.weibo.get(url, new(FriendsTimelineIds))
}

func (self *statuses) UserTimeline(access_token string, uid int64, screenName string, options map[string]interface{}) <-chan Result {
	must := make(map[string]interface{})
	if screenName == "" {
		must["uid"] = uid
	} else {
		must["screen_name"] = screenName
	}
	url := self.weibo.makeUrl("2/statuses/user_timeline.json", access_token, must, options)
	return self.weibo.get(url, new(UserTimeline))
}

func (self *statuses) UserTimelineIds(access_token string, uid int64, screenName string, options map[string]interface{}) <-chan Result {
	must := make(map[string]interface{})
	if screenName == "" {
		must["uid"] = uid
	} else {
		must["screen_name"] = screenName
	}
	url := self.weibo.makeUrl("2/statuses/user_timeline/ids.json", access_token, must, options)
	return self.weibo.get(url, new(UserTimelineIds))
}

func (self *statuses) RepostTimeline(access_token string, id int64, options map[string]interface{}) <-chan Result {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/repost_timeline.json", access_token, must, options)
	return self.weibo.get(url, new(RepostTimeline))
}

func (self *statuses) RepostTimelineIds(access_token string, id int64, options map[string]interface{}) <-chan Result {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/repost_timeline/ids.json", access_token, must, options)
	return self.weibo.get(url, new(RepostTimelineIds))
}

func (self *statuses) RepostByMe(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/repost_by_me.json", access_token, nil, options)
	return self.weibo.get(url, new(RepostByMe))
}

func (self *statuses) Mentions(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/mentions.json", access_token, nil, options)
	return self.weibo.get(url, new(StatusMentions))
}

func (self *statuses) MentionsIds(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/mentions/ids.json", access_token, nil, options)
	return self.weibo.get(url, new(StatusMentionsIds))
}

func (self *statuses) BilateralTimeline(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/bilateral_timeline.json", access_token, nil, options)
	return self.weibo.get(url, new(BilateralTimeline))
}

// more about result?
func (self *statuses) Show(access_token string, id int64, options map[string]interface{}) <-chan Result {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/show.json", access_token, must, options)
	return self.weibo.get(url, new(Status))
}

func (self *statuses) Querymid_One(access_token string, id int64, t int, options map[string]interface{}) <-chan Result {
	must := make(map[string]interface{})
	must["is_batch"] = 0
	must["id"] = id
	must["type"] = t
	if _, ok := options["is_batch"]; ok {
		delete(options, "is_batch")
	}
	url := self.weibo.makeUrl("2/statuses/querymid.json", access_token, must, options)
	return self.weibo.get(url, new(Querymid_One))
}

func (self *statuses) Querymid_Batch(access_token string, ids []int64, t int, options map[string]interface{}) <-chan Result {
	idstr := fmt.Sprintf("%d", ids[0])
	for _, id := range ids[1:] {
		idstr += fmt.Sprintf(",%d", id)
	}
	must := make(map[string]interface{})
	must["is_batch"] = 1
	must["id"] = idstr
	must["type"] = t
	if _, ok := options["is_batch"]; ok {
		delete(options, "is_batch")
	}
	url := self.weibo.makeUrl("2/statuses/querymid.json", access_token, must, options)
	return self.weibo.get(url, new(Querymid_Batch))
}

func (self *statuses) Hot_RepostDaily(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/hot/repost_daily.json", access_token, nil, options)
	return self.weibo.get(url, new(HotRepostDaily))
}

func (self *statuses) Hot_RepostWeekly(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/hot/repost_weekly.json", access_token, nil, options)
	return self.weibo.get(url, new(HotRepostWeekly))
}

func (self *statuses) Hot_CommentsDaily(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/hot/comments_daily.json", access_token, nil, options)
	return self.weibo.get(url, new(HotCommentsDaily))
}

func (self *statuses) Hot_CommentsWeekly(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/hot/comments_weekly.json", access_token, nil, options)
	return self.weibo.get(url, new(HotCommentsWeekly))
}

func (self *statuses) Count(access_token string, ids []int64, options map[string]interface{}) <-chan Result {
	idstr := fmt.Sprintf("%d", ids[0])
	for _, id := range ids[1:] {
		idstr += fmt.Sprintf(",%d", id)
	}
	must := make(map[string]interface{})
	must["ids"] = idstr
	url := self.weibo.makeUrl("2/statuses/count.json", access_token, must, options)
	return self.weibo.get(url, new(Count))
}

func (self *statuses) Repost(access_token string, id int64, options map[string]interface{}) <-chan Result {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/repost.json", access_token, must, options)
	return self.weibo.post(url, "", new(Status))
}

func (self *statuses) Destroy(access_token string, id int64, options map[string]interface{}) <-chan Result {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/destroy.json", access_token, must, options)
	return self.weibo.post(url, "", new(Status))
}

func (self *statuses) Update(access_token string, status string, options map[string]interface{}) <-chan Result {
	must := make(map[string]interface{})
	must["status"] = status
	url := self.weibo.makeUrl("2/statuses/update.json", access_token, must, options)
	return self.weibo.post(url, "", new(Status))
}

func (self *statuses) Upload(access_token string, status string, pic string, options map[string]interface{}) <-chan Result {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	w.WriteField("status", status)
	w.WriteField("pic", pic)
	contentType := w.FormDataContentType()
	w.Close()
	url := self.weibo.makeUrl("2/statuses/upload.json", access_token, nil, options)
	return self.weibo.post(url, contentType, new(Status))
}

func (self *statuses) UploadUrlText(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/statuses/upload_url_text.json", access_token, nil, options)
	return self.weibo.post(url, "", new(Status))
}

func (self *statuses) Emotions(access_token string, options map[string]interface{}) <-chan Result {
	url := self.weibo.makeUrl("2/emotions.json", access_token, nil, options)
	return self.weibo.get(url, new(Emotions))
}
