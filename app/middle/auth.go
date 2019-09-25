package middle

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/lang"
	"hatgo/pkg/util"
	"net/http"
)
const HTTP_TOKEN = "auth_token"

//Socket验证
func SocketAuth(wss *util.Ws, openId string){
	if openId != "hi" {
		wss.SendSelf(lang.CONNECT_FAIL_AUTH,lang.CONNECT_FAIL_AUTH)
		wss.CloseCoon()
	}
}

func Auth(c *gin.Context) {
	authCode := http.StatusUnauthorized
	token := c.GetHeader(HTTP_TOKEN)
	if token == "" {
		util.Output(c, authCode, http.StatusText(authCode))
		c.Abort()
		return
	}
	c.Set("uid", 12)
	//c.GetInt64("uid")
	c.Next()
}
