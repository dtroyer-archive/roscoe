// osc.go - OpenStackClient

package main

import (
    "encoding/json"
    "fmt"
    "log"

    "roscoe/client"
    "roscoe/osclib"
//    "roscoe/server"
)


type ServerResponse struct {
    Servers []interface{}
}

func main() {
    // Get auth values from the environment
    var creds osclib.Creds
    c, err := client.NewClient(creds)
    if err != nil {
        log.Fatal(err)
    }

    // Identity API versions
    osclib.GetVersions(c.Auth)

    // Test authentication
    c.Authenticate()
//    fmt.Printf("token: %s\n", c.Token)
//    fmt.Printf("servcat: %s\n", c.ServCat)

    // list servers
    resp, err := c.Get("compute", "/servers")
    if err != nil {
        // TODO(dtroyer): Handle specific errors
		log.Fatal(err)
    }
    var servers ServerResponse
    err = json.Unmarshal(resp.Body, &servers)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("c: %+v\n\n", servers)
}
