package weibo

import (
	"fmt"
	"mime/multipart"
	"bytes"
)

type statuses struct {
	weibo *Weibo
}

func (self *statuses) PublicTimeline(access_token string, options map[string]interface{}) (*PublicTimeline, error) {
	url := self.weibo.makeUrl("2/statuses/public_timeline.json", access_token, nil, options)
	publicTimeline := new(PublicTimeline)
	return publicTimeline, self.weibo.get(url, publicTimeline)
}

func (self *statuses) FriendsTimeline(access_token string, options map[string]interface{}) (*FriendsTimeline, error) {
	url := self.weibo.makeUrl("2/statuses/friends_timeline.json", access_token, nil, options)
	friendsTimeline := new(FriendsTimeline)
	return friendsTimeline, self.weibo.get(url, friendsTimeline)
}

func (self *statuses) HomeTimeline(access_token string, options map[string]interface{}) (*HomeTimeline, error) {
	url := self.weibo.makeUrl("2/statuses/home_timeline.json", access_token, nil, options)
	homeTimeline := new(HomeTimeline)
	return homeTimeline, self.weibo.get(url, homeTimeline)
}

func (self *statuses) FriendsTimelineIds(access_token string, options map[string]interface{}) (*FriendsTimelineIds, error) {
	url := self.weibo.makeUrl("2/statuses/friends_timeline/ids.json", access_token, nil, options)
	friendsTimelineIds := new(FriendsTimelineIds)
	return friendsTimelineIds, self.weibo.get(url, friendsTimelineIds)
}

func (self *statuses) UserTimeline(access_token string, uid int64, screenName string, options map[string]interface{}) (*UserTimeline, error) {
	must := make(map[string]interface{})
	if screenName == "" {
		must["uid"] = uid
	} else {
		must["screen_name"] = screenName
	}
	url := self.weibo.makeUrl("2/statuses/user_timeline.json", access_token, must, options)
	userTimeline := new(UserTimeline)
	return userTimeline, self.weibo.get(url, userTimeline)
}

func (self *statuses) UserTimelineIds(access_token string, uid int64, screenName string, options map[string]interface{}) (*UserTimelineIds, error) {
	must := make(map[string]interface{})
	if screenName == "" {
		must["uid"] = uid
	} else {
		must["screen_name"] = screenName
	}
	url := self.weibo.makeUrl("2/statuses/user_timeline/ids.json", access_token, must, options)
	userTimelineIds := new(UserTimelineIds)
	return userTimelineIds, self.weibo.get(url, userTimelineIds)
}

func (self *statuses) RepostTimeline(access_token string, id int64, options map[string]interface{}) (*RepostTimeline, error) {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/repost_timeline.json", access_token, must, options)
	repostTimeline := new(RepostTimeline)
	return repostTimeline, self.weibo.get(url, repostTimeline)
}

func (self *statuses) RepostTimelineIds(access_token string, id int64, options map[string]interface{}) (*RepostTimelineIds, error) {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/repost_timeline/ids.json", access_token, must, options)
	repostTimelineIds := new(RepostTimelineIds)
	return repostTimelineIds, self.weibo.get(url, repostTimelineIds)
}

func (self *statuses) RepostByMe(access_token string, options map[string]interface{}) (*RepostByMe, error) {
	url := self.weibo.makeUrl("2/statuses/repost_by_me.json", access_token, nil, options)
	repostByMe := new(RepostByMe)
	return repostByMe, self.weibo.get(url, repostByMe)
}

func (self *statuses) Mentions(access_token string, options map[string]interface{}) (*StatusMentions, error) {
	url := self.weibo.makeUrl("2/statuses/mentions.json", access_token, nil, options)
	statusMentions := new(StatusMentions)
	return statusMentions, self.weibo.get(url, statusMentions)
}

func (self *statuses) MentionsIds(access_token string, options map[string]interface{}) (*StatusMentionsIds, error) {
	url := self.weibo.makeUrl("2/statuses/mentions/ids.json", access_token, nil, options)
	statusMentionsIds := new(StatusMentionsIds)
	return statusMentionsIds, self.weibo.get(url, statusMentionsIds)
}

func (self *statuses) BilateralTimeline(access_token string, options map[string]interface{}) (*BilateralTimeline, error) {
	url := self.weibo.makeUrl("2/statuses/bilateral_timeline.json", access_token, nil, options)
	bilateralTimeline := new(BilateralTimeline)
	return bilateralTimeline, self.weibo.get(url, bilateralTimeline)
}

// more about result?
func (self *statuses) Show(access_token string, id int64, options map[string]interface{}) (*Status, error) {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/show.json", access_token, must, options)
	status := new(Status)
	return status, self.weibo.get(url, status)
}

func (self *statuses) Querymid_One(access_token string, id int64, t int, options map[string]interface{}) (*Querymid_One, error) {
	must := make(map[string]interface{})
	must["is_batch"] = 0
	must["id"] = id
	must["type"] = t
	if _, ok := options["is_batch"]; ok {
		delete(options, "is_batch")
	}
	url := self.weibo.makeUrl("2/statuses/querymid.json", access_token, must, options)
	querymid_One := new(Querymid_One)
	return querymid_One, self.weibo.get(url, querymid_One)
}

func (self *statuses) Querymid_Batch(access_token string, ids []int64, t int, options map[string]interface{}) (*Querymid_Batch, error) {
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
	querymid_Batch := new(Querymid_Batch)
	return querymid_Batch, self.weibo.get(url, querymid_Batch)
}

func (self *statuses) Hot_RepostDaily(access_token string, options map[string]interface{}) (*HotRepostDaily, error) {
	url := self.weibo.makeUrl("2/statuses/hot/repost_daily.json", access_token, nil, options)
	hotRepostDaily := new(HotRepostDaily)
	return hotRepostDaily, self.weibo.get(url, hotRepostDaily)
}

func (self *statuses) Hot_RepostWeekly(access_token string, options map[string]interface{}) (*HotRepostWeekly, error) {
	url := self.weibo.makeUrl("2/statuses/hot/repost_weekly.json", access_token, nil, options)
	hotRepostWeekly := new(HotRepostWeekly)
	return hotRepostWeekly, self.weibo.get(url, hotRepostWeekly)
}

func (self *statuses) Hot_CommentsDaily(access_token string, options map[string]interface{}) (*HotCommentsDaily, error) {
	url := self.weibo.makeUrl("2/statuses/hot/comments_daily.json", access_token, nil, options)
	hotCommentsDaily := new(HotCommentsDaily)
	return hotCommentsDaily, self.weibo.get(url, hotCommentsDaily)
}

func (self *statuses) Hot_CommentsWeekly(access_token string, options map[string]interface{}) (*HotCommentsWeekly, error) {
	url := self.weibo.makeUrl("2/statuses/hot/comments_weekly.json", access_token, nil, options)
	hotCommentsWeekly := new(HotCommentsWeekly)
	return hotCommentsWeekly, self.weibo.get(url, hotCommentsWeekly)
}

func (self *statuses) Count(access_token string, ids []int64, options map[string]interface{}) (*Count, error) {
	idstr := fmt.Sprintf("%d", ids[0])
	for _, id := range ids[1:] {
		idstr += fmt.Sprintf(",%d", id)
	}
	must := make(map[string]interface{})
	must["ids"] = idstr
	url := self.weibo.makeUrl("2/statuses/count.json", access_token, must, options)
	count := new(Count)
	return count, self.weibo.get(url, count)
}

func (self *statuses) Repost(access_token string, id int64, options map[string]interface{}) (*Status, error) {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/repost.json", access_token, must, options)
	status := new(Status)
	return status, self.weibo.post(url, "", status)
}

func (self *statuses) Destroy(access_token string, id int64, options map[string]interface{}) (*Status, error) {
	must := make(map[string]interface{})
	must["id"] = id
	url := self.weibo.makeUrl("2/statuses/destroy.json", access_token, must, options)
	status := new(Status)
	return status, self.weibo.post(url, "", status)
}

func (self *statuses) Update(access_token string, status string, options map[string]interface{}) (*Status, error) {
	must := make(map[string]interface{})
	must["status"] = status
	url := self.weibo.makeUrl("2/statuses/update.json", access_token, must, options)
	statusr := new(Status)
	return statusr, self.weibo.post(url, "", statusr)
}

func (self *statuses) Upload(access_token string, status string, pic string, options map[string]interface{}) (*Status, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	w.WriteField("status", status)
	w.WriteField("pic", pic)
	contentType := w.FormDataContentType()
	w.Close()
	url := self.weibo.makeUrl("2/statuses/upload.json", access_token, nil, options)
	statusr := new(Status)
	return statusr, self.weibo.post(url, contentType, statusr)
}

func (self *statuses) UploadUrlText(access_token string, options map[string]interface{}) (*Status, error) {
	url := self.weibo.makeUrl("2/statuses/upload_url_text.json", access_token, nil, options)
	status := new(Status)
	return status, self.weibo.post(url, "", status)
}

func (self *statuses) Emotions(access_token string, options map[string]interface{}) (*Emotions, error) {
	url := self.weibo.makeUrl("2/emotions.json", access_token, nil, options)
	emotions := new(Emotions)
	return emotions, self.weibo.get(url, emotions)
}
