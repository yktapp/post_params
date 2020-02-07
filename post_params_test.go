package post_param

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetInt(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("int=8"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetInt(r, "int")
	assert.Equal(t, 8, i, "they should be equal")
}

func TestGetArrayInt(t *testing.T) {
	// correct
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("int=8&int=4"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetArrayInt(r, "int")
	assert.Equal(t, []int{8,4}, i, "they should be equal")

	// wrong
	r2 := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("int=worng"))
	r2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i2:= GetArrayInt(r2, "int")
	assert.Equal(t, []int(nil), i2, "they should be equal")
}

func TestGetArrayFloat32(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("float=8&float=4"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetArrayFloat32(r, "float")
	assert.Equal(t, []float32{8,4}, i, "they should be equal")
}

func TestGetArrayFloat64(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("float=8&float=4"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetArrayFloat64(r, "float")
	assert.Equal(t, []float64{8,4}, i, "they should be equal")
}

func TestGetArrayString(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("string=foo&string=bar"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetArrayString(r, "string")
	assert.Equal(t, []string{"foo","bar"}, i, "they should be equal")
}

func TestGetArrayInt64(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("int=8&int=4"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetArrayInt64(r, "int")
	assert.Equal(t, []int64{8,4}, i, "they should be equal")
}

func TestGetArrayBool(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("bool=1&bool=0"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetArrayBool(r, "bool")
	assert.Equal(t, []int64{1,0}, i, "they should be equal")
}

func TestGetString(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("string=hello"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetString(r, "string")
	assert.Equal(t, "hello", i, "they should be equal")
}

func TestGetFloat32(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("float=8"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetFloat32(r, "float")
	assert.Equal(t, float32(8), i, "they should be equal")
}

func TestGetFloat64(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("float=8"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetFloat64(r, "float")
	assert.Equal(t, float64(8), i, "they should be equal")
}

func TestGetInt64(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("int=8"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetInt64(r, "int")
	assert.Equal(t, int64(8), i, "they should be equal")
}

func TestGetBool(t *testing.T) {
	r := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("bool=true"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i:= GetBool(r, "bool")
	assert.Equal(t, true, i, "they should be equal")
	assert.NotEqual(t, false, i, "they should be not equal")

	r2 := httptest.NewRequest("POST", "http://someapi.ru", strings.NewReader("bool=false"))
	r2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	i2:= GetBool(r2, "bool")
	assert.Equal(t, false, i2, "they should be equal")
	assert.NotEqual(t, true, i2, "they should be not equal")
}
