package server

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendString(t *testing.T) {
	w:= testWritter{}
	sut:= "hello"

	err:= SendString(200,sut,&w)

	assert.Nil(t,err)
	assert.Equal(t,sut,w.content)
	assert.Equal(t,200,w.HeaderValue)
}

func TestSendJson(t *testing.T) {
	w:= testWritter{}
	sut:= SerializeTest{
		Id: "1",
		Name: "test",
	}
	expectedString,_ := json.Marshal(sut)
	
	err:= SendJson(200,sut,&w)

	assert.Nil(t,err)
	assert.Equal(t,200,w.HeaderValue)
	assert.Equal(t,string(expectedString),w.content)
}

type SerializeTest struct {
	Id   string
	Name string
}

type testWritter struct {
	HeaderValue int
	content     string
}

func (t *testWritter) Header() http.Header {
	return http.Header{}
}
func (t *testWritter) Write(content []byte) (int, error) {
	t.content = string(content)
	return len(content), nil
}

func (t *testWritter) WriteHeader(statusCode int) {
	t.HeaderValue = statusCode
}
