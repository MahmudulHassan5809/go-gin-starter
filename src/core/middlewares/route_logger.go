package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	Level        string      `json:"level"`
	Status       int         `json:"status"`
	Method       string      `json:"method"`
	Path         string      `json:"path"`
	QueryParams  interface{} `json:"query_params"`
	Headers      interface{} `json:"headers"`
	ReqBody      string      `json:"req_body"`
	ResBody      interface{} `json:"res_body"`
	Latency      string      `json:"latency"`
	Timestamp    string      `json:"timestamp"`
}


type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseBodyWriter) Write(b []byte) (int, error) {
	rw.body.Write(b) // Capture response body
	return rw.ResponseWriter.Write(b)
}


func RequestResponseLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqBody, _ := io.ReadAll(ctx.Request.Body)
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody)) // Restore request body
		resBody := &bytes.Buffer{}
		ginBodyLogger := &responseBodyWriter{body: resBody, ResponseWriter: ctx.Writer}
		ctx.Writer = ginBodyLogger
		startTime := time.Now()
		ctx.Next()
		var parsedResBody interface{}
		if err := json.Unmarshal(resBody.Bytes(), &parsedResBody); err != nil {
			parsedResBody = resBody.String()
		}

		logEntry := LogEntry{
			Status:      ctx.Writer.Status(),
			Method:      ctx.Request.Method,
			Path:        ctx.Request.URL.Path,
			QueryParams: ctx.Request.URL.Query(),
			ReqBody:     string(reqBody),
			ResBody:     parsedResBody,
			Latency:     time.Since(startTime).String(),
			Timestamp:   time.Now().Format(time.RFC3339),
		}
		logJSON, err := json.MarshalIndent(logEntry, "", "    ")
		if err != nil {
			log.Printf("Error marshaling log entry: %v", err)
			return
		}

		fmt.Println(string(logJSON))
	}
}

