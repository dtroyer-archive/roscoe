// flavor/flavor.go - roscoe flavor API package

package flavor

import (
//    "encoding/json"
//    "fmt"
//    "strings"

    "roscoe/client"
    "roscoe/json"
)


type Flavorx struct {
    Id string
    Name string
    Ram string
    Disk string
    Vcpus string
    Links []FlavorLink
    Extra []interface{}
}

type Flavor struct {
    A map[string]interface{}
}

type FlavorLink struct {
    Rel string
    Href string
}

type FlavorResponse json.JsonStruct

type Attr map[string]string

// Compute API v2 Flavors
var FlavorAttrs = []string{"minDisk", "minRam", "marker", "limit"}

var FlavorFields = []string{"id", "name", "ram", "disk", "vcpus", "links"}


// Compute v2 4.4.1: list flavors
func List(c *client.Client, f string) (body *[]interface{}, err error) {
    resp, err := c.Get("compute", "/flavors")
    if err != nil {
        return nil, err
    }

//    var body FlavorResponse
//    err = json.Unmarshal(resp.Body, &body)
    js, err := json.NewJsonStruct(resp.Body)
    if err != nil {
        return nil, err
    }
    b, err := js.GetKey("flavors").Array()
    return &b, nil
}

/*
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

    var dat map[string]interface{}
    err = json.Unmarshal(resp.Body, &dat)
    if err != nil {
        return nil, err
    }
//    fmt.Printf("body: %+v\n", body.Flavors[0].A)
    for k, v := range dat["flavors"].([]interface{}) {
        fmt.Printf("%s=%s\n\n", k, v)
    }
    return nil, err
}
*/