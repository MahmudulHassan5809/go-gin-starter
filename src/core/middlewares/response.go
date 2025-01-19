package middlewares

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

type HttpResponse struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Meta    interface{} `json:"meta"`
	Error   interface{} `json:"error,omitempty"`
}

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		wb := &BodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = wb
		c.Next()
		statusCode := c.Writer.Status()
		rawBody := wb.body.String()
		var responseBody map[string]interface{}
		var responseData interface{}
		var errorResponse interface{}
		if err := json.Unmarshal([]byte(rawBody), &responseBody); err == nil {
			if statusCode != http.StatusOK {
				errorResponse = responseBody["error"]
				responseData = nil
			} else {
				if dataField, exists := responseBody["data"]; exists {
					if nestedData, nestedExists := dataField.(map[string]interface{})["data"]; nestedExists {
						responseData = nestedData
					} else {
						responseData = dataField
					}
				} else {
					responseData = responseBody
				}
			}
		} else {
			log.Println("Response is not in JSON format, passing it through.")
			responseData = rawBody
		}
		finalResponse := HttpResponse{
			Data:    responseData,
			Code:    statusCode,
			Success: (statusCode == http.StatusOK),
			Meta:    nil,
			Error:   errorResponse,
		}
		finalJSON, err := json.Marshal(finalResponse)
		if err != nil {
			log.Printf("Error marshalling final response: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate response"})
			return
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(statusCode)
		c.Writer.WriteString(string(finalJSON))
		wb.body.Reset()
	}
}
