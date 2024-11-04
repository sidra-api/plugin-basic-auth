package main

import (
	"encoding/base64"

	"github.com/sidra-gateway/go-pdk/server"
)

func main() {

	username := "foo"
	password := "bar"

	server.NewServer("basic-auth", func(r server.Request) server.Response {
		//@TODO: Implement your plugin logic here
		if r.Headers["Authorization"] != "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)) {
			return server.Response{
				StatusCode: 401,
				Body:       "Unauthorized",
			}
		}
		return server.Response{
			StatusCode: 200,
			Body:       "Hello, World!",
		}
	})
}
