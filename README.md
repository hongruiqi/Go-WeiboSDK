Go-WeiboSDK
===========

Sina weibo api in golang
-----------

## Example

```go
	wb := weibo.New("APPKey", "APPSecret")
	// options is of type map[string]interface{} for addition params.
	// sync call. For async, please use goroutine outside.
	userTimeline, err := wb.Statuses.UserTimeline(access_token, uid, "", options) 
	if err!=nil {
		panic(err)
	}
	fmt.Println(userTimeline.Statuses)
```
