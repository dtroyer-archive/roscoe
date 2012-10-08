// osc.go - roscoe command line

package main

import (
    "fmt"
    "log"
    "os"

    "github.com/voxelbrain/goptions"

    "roscoe/client"
    "roscoe/flavor"
    "roscoe/osclib"
//    "roscoe/server"
)


func main() {
    options := struct {
        Debug bool      `goptions:"-x, --debug, description='Enable debugging'"`
        Verbose bool    `goptions:"-v, --verbose, description='Be not quiet with output'"`
        goptions.Verbs
        Show struct {
                Server        string `goptions:"-s, --server, description='Server to connect to'"`
        } `goptions:"show"`
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

    print("verb: ", options.Verbs, " ", os.Args[2], "\n")

    // Identity API versions
    osclib.GetVersions(c.Auth)

    // Test authentication
    c.Authenticate()
    if options.Debug {
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
