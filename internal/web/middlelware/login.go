package middlelware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 步骤三
		// 不需要登录的
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		//if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
		//	return
		//}

		sess := sessions.Default(ctx)
		if sess == nil {
			// 没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		id := sess.Get("userId")
		if id == nil {
			// 没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
