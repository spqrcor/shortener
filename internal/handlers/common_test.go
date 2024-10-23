package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_isValidInputParams(t *testing.T) {

	tests := []struct {
		name           string
		method         string
		contentType    string
		reqMethod      string
		reqContentType string
		result         bool
	}{
		{
			name:           "success",
			method:         http.MethodPost,
			contentType:    "application/json",
			reqMethod:      http.MethodPost,
			reqContentType: "application/json",
			result:         true,
		},
		{
			name:           "content type error",
			method:         http.MethodPost,
			contentType:    "text/plain",
			reqMethod:      http.MethodPost,
			reqContentType: "application/json",
			result:         false,
		},
		{
			name:           "method error",
			method:         http.MethodPost,
			contentType:    "text/plain",
			reqMethod:      http.MethodGet,
			reqContentType: "text/plain",
			result:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.reqMethod, "/", nil)
			req.Header.Add("Content-Type", tt.reqContentType)
			result := isValidInputParams(req, inputParams{Method: tt.method, ContentType: tt.contentType})
			assert.Equal(t, result, tt.result)
		})
	}
}
