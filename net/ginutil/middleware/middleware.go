package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/beanscc/rango/net/ginutil/render"
	"github.com/gin-gonic/gin"
)

type CORSConf struct {
	AllowOrigin      []string
	AllowHeaders     []string
	AllowMethods     []string
	ExposeHeaders    []string
	AllowCredentials bool
}

var defaultCORSConf = CORSConf{
	AllowOrigin:      []string{"*"},
	AllowHeaders:     nil,
	AllowMethods:     []string{"*"},
	ExposeHeaders:    nil,
	AllowCredentials: false,
}

// CORS 跨域允许
func CORS(conf *CORSConf) gin.HandlerFunc {
	if conf == nil {
		*conf = defaultCORSConf
	}

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", strings.Join(conf.AllowOrigin, ", "))
		c.Header("Access-Control-Allow-Headers", strings.Join(conf.AllowHeaders, ", "))
		c.Header("Access-Control-Allow-Methods", strings.Join(conf.AllowMethods, ", "))
		c.Header("Access-Control-Expose-Headers", strings.Join(conf.ExposeHeaders, ", "))
		c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(conf.AllowCredentials))
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
		}

		c.Next()
	}
}

func NoRouter() gin.HandlerFunc {
	gin.Default()
	return func(c *gin.Context) {
		render.GinJSON(c, 404, "", nil)
	}
}
