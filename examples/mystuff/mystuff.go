// mystuff.go - roscoe client example

package main

import (
    "fmt"
    "log"

    goopt "github.com/droundy/goopt"

    "roscoe/client"
    "roscoe/flavor"
    "roscoe/osclib"
    "roscoe/server"
)


// Declare options

var debug = goopt.Flag(
    []string{"-x", "--debug"},
    []string{"--nodebug"},
    "Enable debug mode",
    "Disable debug mode",
)

var verbose = goopt.Flag(
    []string{"-v", "--verbose"},
    []string{"-q", "--quiet"},
    "output verbosely",
    "be quiet, instead",
)


func main() {
    goopt.Description = func() string {
        return "OpenStack client example"
    }
    goopt.Version = "1.0"
    goopt.Parse(nil)

    // Propagate debug setting to packages
    client.Debug = debug

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
        if *verbose == true {
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
