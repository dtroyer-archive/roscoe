// mystuff.go - roscoe client example

package main

import (
    "fmt"
    "log"

    "github.com/voxelbrain/goptions"

    "roscoe/client"
    "roscoe/flavor"
    "roscoe/osclib"
    "roscoe/server"
)


func main() {
    options := struct {
        Debug bool      `goptions:"-x, --debug, description='Enable debugging'"`
        Verbose bool    `goptions:"-v, --verbose, description='Be not quiet with output'"`
    }{
        Debug: false,
        Verbose: false,
    }
    goptions.Parse(&options)

    // Propagate debug setting to packages
    client.Debug = &options.Debug

    // Get auth values from the environment
    var creds osclib.Creds
    c, err := client.NewClient(creds)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print("Servers: \n")
    servers, err := server.List(c, "")
    if err != nil {
        log.Fatal(err)
    }
    if servers != nil {
        if options.Verbose == true {
            fmt.Printf("c: %+v\n\n", *servers)
        } else {
            fmt.Printf("c: %+v\n\n", *servers)
        }
    }

    flavors, err := flavor.List(c, "")
    if err != nil {
        log.Fatal(err)
    }
//    osclib.OutputData([]string{"a","b","c"}, []interface{}(flavors))
    for k, v := range flavors.Flavors {
        fmt.Printf("%d: %s\n", k, v.Id)
    }

}
