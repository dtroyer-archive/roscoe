// server.go

package server

import (
    "encoding/json"

    "roscoe/client"
)

type Server struct {
    Id string
    Name string
}


type ServerResponse struct {
    Servers []interface{}
}


func List(c *client.Client, f string) (servers *ServerResponse, err error) {
    // list servers
    resp, err := c.Get("compute", "/servers")
    if err != nil {
        return nil, err
    }

    var s ServerResponse
    err = json.Unmarshal(resp.Body, &s)
    if err != nil {
        return nil, err
    }

    return &s, nil
}