package weibo

type account struct {
	weibo *Weibo
}

func (self *account) GetUid(access_token string) (*Uid, error) {
	url := self.weibo.makeUrl("2/account/get_uid.json", access_token, nil, nil)
	uid := new(Uid)
	return uid, self.weibo.get(url, uid)
}
