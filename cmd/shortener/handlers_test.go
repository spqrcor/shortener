package main

import (
	"net/http"
	"testing"
)

func Test_createShortHandler(t *testing.T) {
	type args struct {
		res http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createShortHandler(tt.args.res, tt.args.req)
		})
	}
}

func Test_searchShortHandler(t *testing.T) {
	type args struct {
		res http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searchShortHandler(tt.args.res, tt.args.req)
		})
	}
}
