package handlers

import (
	"net/http"
	"strings"
)

type inputJSONData struct {
	URL string `json:"url,omitempty"`
}

type outputJSONData struct {
	Result string `json:"result,omitempty"`
}

type inputParams struct {
	Method      string
	ContentType string
}

func isValidInputParams(req *http.Request, params inputParams) bool {
	if req.Method != params.Method {
		return false
	}
	if req.Header.Get(`Content-Encoding`) != `gzip` && !strings.Contains(req.Header.Get("Content-Type"), params.ContentType) {
		return false
	}
	return true
}
