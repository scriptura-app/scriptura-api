package utils

import (
	"net/http"
	"strconv"
)

func GetParam(r *http.Request, param string, defaultValue string) string {
	valueStr := r.URL.Query().Get(param)
	if valueStr != "" {
		return valueStr
	}
	return defaultValue
}

func GetIntParam(r *http.Request, param string, defaultValue int) int {
	valueStr := r.URL.Query().Get(param)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
