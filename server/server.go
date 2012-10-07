// server.go - roscoe server API package

package server

import (
    "encoding/json"
    "strings"

    "roscoe/client"
)


type Server struct {
    Id string
    Name string
}


type ServerResponse struct {
    Servers []interface{}
}

type Attr map[string]string

// Compute API v2 Servers
var ServerAttrs = []string{"name", "image", "flavor", "status", "marker", "limit", "changes-since"}


// Compute v2 4.1.1: list servers
func List(c *client.Client, f string) (servers *ServerResponse, err error) {
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

// Compute v2 4.1.1: list servers
func Show(c *client.Client, attr Attr) (s *ServerResponse, err error) {
    print("attr: ", attr["name"], "\n")
    // Look for search filters
    var f []string
    for _, v := range ServerAttrs {
        if val, ok := attr[v]; ok {
            f = append(f, attr[val])
        }
    }
    filter := strings.Join(f, "&")
    print("filter: ", filter, "\n")

    resp, err := c.Get("compute", "/servers/detail")
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(resp.Body, &s)
    if err != nil {
        return nil, err
    }

    return
}
