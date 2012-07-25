Go-WeiboSDK
===========

Sina weibo api in golang
-----------

## Example

```go
	wb := weibo.New("APPKey", "APPSecret")
	
	// errChan is the waiting channel for the result.
	// userTimeline can't be read until errChan received nil.
	// A received error indicates that something wrong in api call.
	// Options is of type map[string]interface{} for addition params.
	userTimeline, errChan := wb.Statuses.UserTimeline(access_token, uid, "", options) 
	if err:=<-errChan; err!=nil {
		panic(err)
	}
	fmt.Println(userTimeline.Statuses)
```
