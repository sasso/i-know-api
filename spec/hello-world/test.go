package hello_world

import (
	"testing"

	"gopkg.in/h2non/baloo.v1"
)

// test stores the HTTP testing client preconfigured
var test = baloo.New("http://httpbin.org")

func TestBalooBasic(t *testing.T) {
	test.Get("/get").
		SetHeader("Foo", "Bar").
		Expect(t).
		Status(200).
		Header("Server", "nginx").
		Type("json").
		Done()
}
