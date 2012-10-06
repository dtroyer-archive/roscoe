// osc.go - OpenStackClient

package main

import (
    "flag"
    "fmt"
    "log"

    "roscoe/client"
    "roscoe/osclib"
    "roscoe/server"
)

// debug flag
var debug = flag.Bool("x", false, "set debug mode")

func main() {
    flag.Parse()
    client.Debug = debug

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
    if *debug {
        fmt.Printf("token: %s\n", c.Token)
        fmt.Printf("servcat: %s\n", c.ServCat)
    }

    servers, err := server.List(c, "")
    fmt.Printf("c: %+v\n\n", *servers)
}
