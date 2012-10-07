// osc.go - roscoe command line

package main

import (
    "flag"
    "fmt"
    "log"

    "roscoe/client"
    "roscoe/flavor"
    "roscoe/osclib"
//    "roscoe/server"
)


// debug flag
var debug = flag.Bool("x", false, "Enable debug mode")


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

//    servers, err := server.List(c, "")
//    fmt.Printf("c: %+v\n\n", *servers)

//    attr := make(server.Attr)
//    attr["name"] = "npd01"
//    servers, err := server.Show(c, attr)
//    if err != nil {
//        log.Fatal(err)
//    }
//    fmt.Printf("c: %+v\n\n", *servers)

    flavors, err := flavor.List(c, "")
    if err != nil {
        log.Fatal(err)
    }
    for k, v := range flavors.Flavors {
        fmt.Printf("%d: %s\n", k, v.Id)
    }
}
