package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	gsd "github.com/yylt/gorilla-session-django"
)

var (
	sename         = "k"
	defaultId      = "sessionid"
	defaultPrefix  = ":1:django.contrib.sessions.cache"
	sessionAuthKey = "_auth_user_id"

	sestore *gsd.Store
)

func SessionAuthMid(memcli gsd.Memcacher) gin.HandlerFunc {
	var (
		gsdcfg = &gsd.Gsd_config{
			CookieKey:      defaultId,
			Auth:           nil,
			JsonSerializer: false,
			PrefixMemcache: defaultPrefix,
		}
	)
	gsdcfg.Auth = func(value map[string]string) error {
		_, ok := value[sessionAuthKey]
		if ok {
			return nil
		}
		return fmt.Errorf("Not found authKey in session")
	}
	sestore = gsd.NewSessionStore(memcli, gsdcfg, nil)

	return func(context *gin.Context) {
		_, err := sessions.GetRegistry(context.Request).Get(sestore, sename)
		if err != nil {
			context.AbortWithStatusJSON(401, map[string]string{"error": err.Error()})
			return
		}
		context.Next()
	}
}

func GetSessionValue(key string) (string, error) {
	dict := sestore.Values()
	if dict == nil {
		return "", fmt.Errorf("Session is null")
	}
	v, ok := dict[key]
	if ok {
		return v, nil
	}
	return "", fmt.Errorf("Not found key:%s", key)
}
