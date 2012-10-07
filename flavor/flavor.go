// flavor/flavor.go - roscoe flavor API package

package flavor

import (
    "encoding/json"
    "strings"

    "roscoe/client"
)


type Flavor struct {
    Id string
    Name string
    Ram string
    Disk string
    Vcpus string
    Links []FlavorLink
    Extra []interface{}
}

type FlavorLink struct {
    Rel string
    Href string
}

type FlavorResponse struct {
    Flavors []Flavor
}

type Attr map[string]string

// Compute API v2 Flavors
var FlavorAttrs = []string{"minDisk", "minRam", "marker", "limit"}

var FlavorFields = []string{"id", "name", "ram", "disk", "vcpus", "links"}


// Compute v2 4.4.1: list flavors
func List(c *client.Client, f string) (body *FlavorResponse, err error) {
    resp, err := c.Get("compute", "/flavors")
    if err != nil {
        return nil, err
    }

//    var body FlavorResponse
    err = json.Unmarshal(resp.Body, &body)
    if err != nil {
        return nil, err
    }

    return body, nil
}

// Compute v2 4.4.1: list flavor details
func Show(c *client.Client, attr Attr) (body *FlavorResponse, err error) {
    // Look for search filters
    var f []string
    print("attr[1]: ", FlavorAttrs[1],"\n")
    for _, v := range FlavorAttrs {
    print("v: ",v,"\n")
        if val, ok := attr[v]; ok {
            f = append(f, attr[val])
        }
    }
    filter := strings.Join(f, "&")
    print("filter: ", filter, "\n")

    resp, err := c.Get("compute", "/flavors/detail")
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(resp.Body, &body)
    if err != nil {
        return nil, err
    }

    return
}
