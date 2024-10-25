package messageFormat


type Message map[string]string

var MsgFormat Message

func(m Message)Success(data string) map[string]string {
    return map[string]string{"errorCode":"0", "message":"successful", "data": data}
}

func(m Message)Failed() map[string]string {
    return map[string]string{"errorCode":"1", "message":"params error"}
}

func(m Message)DBFailed() map[string]string {
    return map[string]string{"errorCode":"2", "message":"Database error"}
}

func(m Message)TokenFailed() map[string]string {
    return map[string]string{"errorCode":"3", "message":"token error"}
}


func(m Message)TokenLose() map[string]string {
    return map[string]string{"errorCode":"4", "message":"miss token"}
}


func(m Message)CustomMessage(msg string) map[string]string {
    return map[string]string{"errorCode":"99", "message":msg}
}