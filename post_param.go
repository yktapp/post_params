package post_param

import (
	"net/http"
	"strconv"
)

func GetInt(r *http.Request, key string) (value int) {
	val := r.PostFormValue(key)
	value, _ = strconv.Atoi(val)
	return value
}

func GetInt64(r *http.Request, key string) (value int64) {
	val := r.PostFormValue(key)
	value, _ = strconv.ParseInt(val, 10, 64)
	return value
}

func GetFloat32(r *http.Request, key string) (value float32) {
	val := r.PostFormValue(key)
	value64, _ := strconv.ParseFloat(val, 32)
	return float32(value64)
}

func GetFloat64(r *http.Request, key string) (value float64) {
	val := r.PostFormValue(key)
	value, _ = strconv.ParseFloat(val, 32)
	return value
}

func GetBool(r *http.Request, key string) (value bool) {
	val := r.PostFormValue(key)
	value, _ = strconv.ParseBool(val)
	return value
}

func GetString(r *http.Request, key string) string {
	return r.PostFormValue(key)
}
