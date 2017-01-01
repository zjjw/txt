package api_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	"github.com/zjjw/txt/api"
	"github.com/zjjw/txt/models"
	"net/http/httptest"
	"testing"
)

func TestGetPost(t *testing.T) {
	ps := httprouter.Params{
		httprouter.Param{"id", "1"},
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	api.GetPost(w, req, ps)

	t.Logf("%d - %s", w.Code, w.Body.String())

	if w.Code != 200 {
		t.Fail()
	}

	post := &models.Post{}
	err := proto.Unmarshal(w.Body.Bytes(), post)

	if err != nil {
		t.Fail()
	}

	if post.Id != "1" {
		t.Fail()
	}

	if post.Contents != "hello world" {
		t.Fail()
	}
}

func TestGetPosts(t *testing.T) {
	ps := httprouter.Params{}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	api.GetPosts(w, req, ps)

	t.Logf("%d - %s", w.Code, w.Body.String())

	if w.Code != 200 {
		t.Fail()
	}

	posts := &models.Posts{}
	err := proto.Unmarshal(w.Body.Bytes(), posts)

	if err != nil {
		t.Fail()
	}

	post := posts.Data[0]

	if post.Id != "1" {
		t.Fail()
	}

	if post.Contents != "hello world" {
		t.Fail()
	}
}
