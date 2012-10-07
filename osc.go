// osc.go - roscoe command line

package main

import (
    "fmt"
    "log"

    goopt "github.com/droundy/goopt"

    "roscoe/client"
    "roscoe/flavor"
    "roscoe/osclib"
//    "roscoe/server"
)


// debug option
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
        return "OpenStack client"
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
