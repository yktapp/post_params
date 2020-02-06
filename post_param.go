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

func GetArrayInt(r *http.Request, key string) (values []int) {
	if err := r.ParseForm(); err != nil{
		return nil
	}
	for _, s := range r.Form[key]{
		val, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		values = append(values, val)
	}
	return values
}

func GetArrayInt64(r *http.Request, key string) (values []int64) {
	if err := r.ParseForm(); err != nil{
		return nil
	}
	for _, s := range r.Form[key]{
		val, err := strconv.ParseInt(s,10,64)
		if err != nil {
			continue
		}
		values = append(values, val)
	}
	return values
}

func GetArrayBool(r *http.Request, key string) (values []int64) {
	if err := r.ParseForm(); err != nil{
		return nil
	}
	for _, s := range r.Form[key]{
		val, err := strconv.ParseInt(s,10,64)
		if err != nil {
			continue
		}
		values = append(values, val)
	}
	return values
}

func GetArrayFloat32(r *http.Request, key string) (values []float32) {
	if err := r.ParseForm(); err != nil{
		return nil
	}
	for _, s := range r.Form[key]{
		val, err := strconv.ParseFloat(s,32)
		if err != nil {
			continue
		}
		values = append(values, float32(val))
	}
	return values
}

func GetArrayFloat64(r *http.Request, key string) (values []float64) {
	if err := r.ParseForm(); err != nil{
		return nil
	}
	for _, s := range r.Form[key]{
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			continue
		}
		values = append(values, val)
	}
	return values
}