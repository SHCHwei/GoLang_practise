package responseFormat


type Message struct{}

var MsgFormat Message

func(m Message)Success(data interface{}) map[string]interface{} {
    temp := make(map[string]interface{})
    temp["errorCode"] = "0"
    temp["message"] = "successful"
    temp["data"] = data
    return temp
}

func(m Message)Failed(data interface{}) map[string]interface{} {
    temp := make(map[string]interface{})
    temp["errorCode"] = "1"
    temp["message"] = "params error"
    temp["system"] = data
    return temp
}

func(m Message)DBFailed(data interface{}) map[string]interface{} {
    temp := make(map[string]interface{})
    temp["errorCode"] = "2"
    temp["message"] = "Database error"
    temp["systemMessage"] = data
    return temp
}

func(m Message)TokenFailed() map[string]interface{} {
    temp := make(map[string]interface{})
    temp["errorCode"] = "3"
    temp["message"] = "token error"
    return temp
}


func(m Message)TokenLose() map[string]interface{} {
    temp := make(map[string]interface{})
    temp["errorCode"] = "4"
    temp["message"] = "miss token"
    return temp
}


func(m Message)CustomMessage(data interface{}) map[string]interface{} {
    temp := make(map[string]interface{})
    temp["errorCode"] = "99"
    temp["message"] = data
    return temp
}