package middleware

import(
    "github.com/gin-gonic/gin"
    Rf "crud/pkg/responseFormat"
    "crypto/sha256"
    "reflect"
    "encoding/hex"
    "log"
)

const apiToken string = "aor68jhty38mutd9"

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {

        var checked bool = false

        // Step1 檢查 Token & Unix time
        unixTime := reflect.ValueOf(c.PostForm("Unix"))

        if c.PostForm("Token") == "" || unixTime.Len() == 0 {
            c.JSON(200, Rf.MsgFormat.TokenLose())
	        c.Abort()
        }else{

            // Step2 unix & apiToken 做出 16進位[]byte token
            token16ByteArr := sha256.Sum256([]byte(unixTime.String() + apiToken))

            token, err := hex.DecodeString(c.PostForm("Token"))

            if err != nil {
                c.JSON(200, Rf.MsgFormat.TokenFailed())
            	c.Abort()
            }

            // Step 3 token16ByteArr 和 Token 比對
            if len(token16ByteArr) != len(token) {
                c.JSON(200, Rf.MsgFormat.TokenLose())
                c.Abort()
            }else{
                for k, v := range token16ByteArr{
                    if v != token[k] {
                        checked = true
                        log.Println("error Point => ", k , v , token[k])
                    }
                }
            }

            if checked {
                c.JSON(200, Rf.MsgFormat.TokenFailed())
            	c.Abort()
            }else{
                c.Next()
            }
        }
	}
}