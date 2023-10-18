package render

import (
	"net/http"
	"reflect"

	"github.com/beanscc/rango/ecode"
	"github.com/gin-gonic/gin"
)

var (
	codeText      = map[int]string{}
	codeOk        int
	codeSystemErr int
)

type JSONResp struct {
	Code int         `json:"retCode"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func RegisterCode(code map[int]string, ok, systemErr int) {
	codeText = code
	codeOk = ok
	codeSystemErr = systemErr
}

// GinJSON gin json 响应
// 使用 GinJSON(c, 404, "tip msg", data)
func GinJSON(c *gin.Context, code int, msg string, data interface{}) {
	if data != nil {
		// 处理 data 值为 nil 时，json 序列化后 null 的问题
		rv := reflect.ValueOf(data)
		if rv.IsNil() {
			data = nil
		}
	}

	if msg == "" {
		msg = codeText[code]
	}

	res := JSONResp{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	c.Set("resp", &res)
	c.Set("resp.code", code)
	c.JSON(http.StatusOK, &res)
}

func GinJSONCode(c *gin.Context, code int) {
	GinJSON(c, code, "", nil)
}

func GinJSONErr(c *gin.Context, err error) {
	e1, ok := ecode.FromError(err)
	if ok {
		GinJSON(c, e1.Code(), e1.Msg(), nil)
		return
	}
	GinJSON(c, codeSystemErr, "", nil)
}

func GinJSONOk(c *gin.Context, data interface{}) {
	GinJSON(c, codeOk, "", data)
}

/*

register code
register render
// 成功响应
Success(ctx context.Context, data interface{})
Failed(ctx context.Context, err error)
Render(ctx context.Context, code int)


选择 render, 设置 data to obj 的func
*/

type Render interface {
	// Render(ctx *gin.Context, httpStatusCode int, obj interface{})
	Success(ctx *gin.Context, data interface{})
	Failed(ctx *gin.Context, data interface{})
	Msg(ctx *gin.Context, data interface{})
}

type JSON struct {
	ToObjFn func(data interface{}, err error) (obj interface{})
}

func (r *JSON) Render(ctx *gin.Context, httpStatus int, obj interface{}) {
	ctx.JSON(httpStatus, obj)
}

// ToObj 根据 data 和 err 构建 返回数据
func (r *JSON) ToObj(data interface{}, err error) interface{} {
	return r.ToObjFn(data, err)
}

// Success
func (r *JSON) Success(ctx *gin.Context, data interface{}) {
	r.Render(ctx, http.StatusOK, r.ToObj(data, nil))
}

// Failed
func (r *JSON) Failed(ctx *gin.Context, err error) {
	r.Render(ctx, http.StatusOK, r.ToObj(nil, err))
}
