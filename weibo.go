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

type Result struct {
	Reply APIReply
	Error error
}

type Weibo struct {
	clientId     string
	clientSecret string
	client       *http.Client
	Statuses     *statuses
}

func New(clientId, clientSecret string) *Weibo {
	tr := &http.Transport{MaxIdleConnsPerHost: 5}
	client := &http.Client{Transport: tr}
	weibo := &Weibo{clientId: clientId, clientSecret: clientSecret, client: client}
	weibo.Statuses = &statuses{weibo: weibo}
	return weibo
}

func (weibo *Weibo) AccessToken(code string, redirectUri string) <-chan Result {
	paramsFmt := "client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s&code=%s"
	params := fmt.Sprintf(paramsFmt,
		weibo.clientId,
		weibo.clientSecret,
		redirectUri,
		code)
	url := fmt.Sprintf("%s%s?%s", WEIBO_OAUTH2_URL, "access_token", params)
	return weibo.post(url, "", new(AccessToken))
}

func (weibo *Weibo) get(url string, reply APIReply) <-chan Result {
	resultChan := make(chan Result, 2)
	go weibo.call(url, true, "", reply, resultChan)
	return resultChan
}

func (weibo *Weibo) post(url string, contentType string, reply APIReply) <-chan Result {
	if contentType == "" {
		contentType = "application/x-www-form-encoded"
	}
	resultChan := make(chan Result, 2)
	weibo.call(url, false, contentType, reply, resultChan)
	return resultChan
}

func (weibo *Weibo) call(url string, get bool, contentType string, reply APIReply, resultChan chan Result) {
	log.Println(url)
	var resp *http.Response
	var err error
	if get {
		resp, err = weibo.client.Get(url)
	} else {
		resp, err = weibo.client.Post(url, contentType, nil)
	}
	if err != nil {
		resultChan <- Result{Error: err}
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode == 200 {
		err = dec.Decode(reply)
		if err != nil {
			resultChan <- Result{Error: err}
			return
		}
		resultChan <- Result{Reply: reply}
		return
	}
	APIErr := new(APIError)
	err = dec.Decode(APIErr)
	if err != nil {
		resultChan <- Result{Error: err}
		return
	}
	resultChan <- Result{Error: APIErr}
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
