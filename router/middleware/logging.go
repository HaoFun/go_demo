package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/willf/pad"

	"go_demo/handler"
	"go_demo/package/errno"
)

type bodyLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWrite) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now().UTC()
		path := ctx.Request.URL.Path

		reg := regexp.MustCompile("(/v1/user|/login)")
		if !reg.MatchString(path) {
			return
		}

		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
		}

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		method := ctx.Request.Method
		ip := ctx.ClientIP()

		bodyLogWriter := &bodyLogWrite{
			body: bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}

		ctx.Writer = bodyLogWriter
		ctx.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		var response handler.Response
		if err := json.Unmarshal(bodyLogWriter.body.Bytes(), &response); err != nil {
			log.Errorf(
				err,
				"response body can not unmarshal to model. Response struct, body: `%s`",
				bodyLogWriter.body.Bytes(),
			)
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}

		log.Infof(
			"%-13s | %-12s | %s %s | {code: %d, message: %s}",
			latency,
			ip,
			pad.Right(method, 5, ""),
			path,
			code,
			message,
		)
	}
}