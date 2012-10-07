// mystuff.go - roscoe client example

package main

import (
    "flag"
    "fmt"
    "log"

    "roscoe/client"
    "roscoe/flavor"
    "roscoe/osclib"
    "roscoe/server"
)


// debug flag
var debug = flag.Bool("x", false, "Enable debug mode")


func main() {
    help := flag.Bool("help", false, "Show usage")
    verbose := flag.Bool("v", false, "Show token details")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
    client.Debug = debug

    if *help == true {
        flag.Usage()
        return
    }

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
