package common

import "testing"

type TestGetJSONResp struct {
	Data struct {
		Key string `json:"key"`
	} `json:"args"`
}

type TestPostJSONReq struct {
	Key string `json:"key"`
}

type TestPostJSONResp struct {
	Data struct {
		Key string `json:"key"`
	} `json:"json"`
}

func TestGetJSON(t *testing.T) {
	var err error
	var resp TestGetJSONResp
	if err = GetJSON("http://httpbin.org/get?key=val", &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Data.Key != "val" {
		t.Fatal("not equal")
	}
}

func TestPostJSON(t *testing.T) {
	var err error
	var req TestPostJSONReq
	req.Key = "val"
	var resp TestPostJSONResp
	if err = PostJSON("http://httpbin.org/post", &req, nil); err != nil {
		t.Fatal(err)
	}
	if err = PostJSON("http://httpbin.org/post", &req, &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Data.Key != "val" {
		t.Fatal("not equal")
	}
}

func TestPutJSON(t *testing.T) {
	var err error
	var req TestPostJSONReq
	req.Key = "val"
	var resp TestPostJSONResp
	if err = PutJSON("http://httpbin.org/put", &req, nil); err != nil {
		t.Fatal(err)
	}
	if err = PutJSON("http://httpbin.org/put", &req, &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Data.Key != "val" {
		t.Fatal("not equal")
	}
}

func TestPatchJSON(t *testing.T) {
	var err error
	var req TestPostJSONReq
	req.Key = "val"
	var resp TestPostJSONResp
	if err = PatchJSON("http://httpbin.org/patch", &req, nil); err != nil {
		t.Fatal(err)
	}
	if err = PatchJSON("http://httpbin.org/patch", &req, &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Data.Key != "val" {
		t.Fatal("not equal")
	}
}

func TestDeleteJSON(t *testing.T) {
	var err error
	var resp TestGetJSONResp
	if err = DeleteJSON("http://httpbin.org/delete?key=val", &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Data.Key != "val" {
		t.Fatal("not equal")
	}
}
