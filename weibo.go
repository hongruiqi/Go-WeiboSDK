package weibo

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

const (
	WEIBO_OAUTH2_URL   = "https://api.weibo.com/oauth2/"
	WEIBO_API_BASE_URL = "https://api.weibo.com/"
)

type APIReply interface{}

type Weibo struct {
	clientId     string
	clientSecret string
	debug        bool
	Statuses     *statuses
	Account      *account
}

func New(clientId, clientSecret string, debug bool) *Weibo {
	weibo := &Weibo{clientId: clientId, clientSecret: clientSecret, debug: debug}
	weibo.Statuses = &statuses{weibo: weibo}
	weibo.Account = &account{weibo: weibo}
	return weibo
}

func (weibo *Weibo) AccessToken(code string, redirectUri string) (*AccessToken, error) {
	paramsFmt := "client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s&code=%s"
	params := fmt.Sprintf(paramsFmt,
		weibo.clientId,
		weibo.clientSecret,
		redirectUri,
		code)
	url := fmt.Sprintf("%s%s?%s", WEIBO_OAUTH2_URL, "access_token", params)
	accessToken := new(AccessToken)
	return accessToken, weibo.post(url, "", accessToken)
}

func (weibo *Weibo) get(url string, reply APIReply) error {
	return weibo.call(url, true, "", reply)
}

func (weibo *Weibo) post(url string, contentType string, reply APIReply) error {
	if contentType == "" {
		contentType = "application/x-www-form-encoded"
	}
	return weibo.call(url, false, contentType, reply)
}

func (weibo *Weibo) call(url string, get bool, contentType string, reply APIReply) (err error) {
	if weibo.debug {
		log.Printf("[WeiboSDK] %s", url)
	}
	var resp *http.Response
	if get {
		resp, err = http.Get(url)
	} else {
		resp, err = http.Post(url, contentType, nil)
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode == 200 {
		err = dec.Decode(reply)
		if err != nil {
			return
		}
		return
	}
	APIErr := new(APIError)
	err = dec.Decode(APIErr)
	if err != nil {
		return
	}
	return
}

func (weibo *Weibo) makeUrl(api string, access_token string, must map[string]interface{}, options map[string]interface{}) string {
	params := fmt.Sprintf("access_token=%s", access_token)
	for k, v := range must {
		params += fmt.Sprintf("&%s=%v", k, v)
	}
	for k, v := range options {
		params += fmt.Sprintf("&%s=%v", k, v)
	}
	return fmt.Sprintf("%s%s?%s", WEIBO_API_BASE_URL, api, params)
}

func (weibo *Weibo) makeUrlSource(api string, must map[string]interface{}, options map[string]interface{}) string {
	params := fmt.Sprintf("source=%s", weibo.clientId)
	for k, v := range must {
		params += fmt.Sprintf("&%s=%v", k, v)
	}
	for k, v := range options {
		params += fmt.Sprintf("&%s=%v", k, v)
	}
	return fmt.Sprintf("%s%s?%s", WEIBO_API_BASE_URL, api, params)
}
