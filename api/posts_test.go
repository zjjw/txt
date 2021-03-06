package api

import (
	"bytes"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
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

	GetPost(w, req, ps)

	t.Logf("%d - %s", w.Code, w.Body.String())

	if w.Code != 200 {
		t.Fail()
	}

	post := &models.Post{}
	err := proto.Unmarshal(w.Body.Bytes(), post)

	if err != nil {
		t.Fail()
	}

	if post.GetId() != "1" {
		t.Fail()
	}

	if post.GetContents().GetBlocks()[0].GetKey() != models.HelloWorldPost.GetContents().GetBlocks()[0].GetKey() {
		t.Fail()
	}
}

func TestGetPosts(t *testing.T) {
	ps := httprouter.Params{}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	GetPosts(w, req, ps)

	t.Logf("%d - %s", w.Code, w.Body.String())

	if w.Code != 200 {
		t.Fail()
	}

	posts := &models.Posts{}
	err := proto.Unmarshal(w.Body.Bytes(), posts)

	if err != nil {
		t.Fail()
	}

	post := posts.Posts[0]

	if post.GetId() != models.HelloWorldPost.GetId() {
		t.Fail()
	}

	if post.GetContents().GetBlocks()[0].GetKey() != models.HelloWorldPost.GetContents().GetBlocks()[0].GetKey() {
		t.Fail()
	}
}

func TestNewPost(t *testing.T) {
	ps := httprouter.Params{}

	data, marshalErr := proto.Marshal(models.HelloWorldPost)
	if marshalErr != nil {
		t.Fail()
	}

	r := bytes.NewReader(data)

	req := httptest.NewRequest("POST", "http://example.com/foo", r)
	w := httptest.NewRecorder()

	hub := NewHub()
	go hub.Run()

	ctx := context.WithValue(req.Context(), "hub", hub)

	// Create the post
	go func() { NewPost(w, req.WithContext(ctx), ps) }()

	// Wait for the created message
	message := <-hub.created

	post := &models.Post{}
	err := proto.Unmarshal(message, post)
	if marshalErr != nil {
		t.Fail()
	}

	if post.GetId() == models.HelloWorldPost.GetId() {
		t.Fail()
	}

	if post.GetContents().GetBlocks()[0].GetKey() != models.HelloWorldPost.GetContents().GetBlocks()[0].GetKey() {
		t.Fail()
	}

	// Assert that the response is good
	t.Logf("%d - %s", w.Code, w.Body.String())

	if w.Code != 200 {
		t.Fail()
	}

	post = &models.Post{}
	err = proto.Unmarshal(w.Body.Bytes(), post)

	if err != nil {
		t.Fail()
	}

	if post.GetId() == models.HelloWorldPost.GetId() {
		t.Fail()
	}

	if post.GetContents().GetBlocks()[0].GetKey() != models.HelloWorldPost.GetContents().GetBlocks()[0].GetKey() {
		t.Fail()
	}
}
