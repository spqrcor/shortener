// Package handlers обработчики http запросов
package handlers

import (
	"net/http"
	"strings"
)

// inputJSONData тип входящего запроса /api/shorten
type inputJSONData struct {
	URL string `json:"url,omitempty"`
}

// outputJSONData тип ответа на запрос /api/shorten
type outputJSONData struct {
	Result string `json:"result,omitempty"`
}

// inputParams тип для валидации http запроса
type inputParams struct {
	Method      string
	ContentType string
}

// isValidInputParams валидация http запроса, req *http.Request, params - параметры запроса
func isValidInputParams(req *http.Request, params inputParams) bool {
	if req.Method != params.Method {
		return false
	}
	if req.Header.Get(`Content-Encoding`) != `gzip` && !strings.Contains(req.Header.Get("Content-Type"), params.ContentType) {
		return false
	}
	return true
}
