// osc.go - roscoe command line

package main

import (
    "errors"
    "fmt"
    "log"
    "os"
    "reflect"

    "github.com/voxelbrain/goptions"

    "roscoe/client"
    "roscoe/flavor"
    //    "roscoe/osclib"
    "roscoe/server"
)

const (
    Version = "0.1"
)

type CmdFunc func(c *client.Client, opts interface{}) error

var ListCmds map[string]CmdFunc = map[string]CmdFunc{
    "flavors": DoListFlavors,
    "servers": DoListServers,
}

var ShowCmds map[string]CmdFunc = map[string]CmdFunc{
    //    "flavor": DoShowFlavor,
    "server": DoShowServer,
}

func getField(opts interface{}, name string) (r reflect.Value) {
    v := reflect.ValueOf(opts).Elem()
    if v.Kind() == reflect.Struct {
        r = v.FieldByName(name)
        return r
    }
    return r
}

func DoListFlavors(c *client.Client, opts interface{}) error {
    if getField(opts, "All").Bool() {
        fmt.Printf("all\n")
    }
    flavors, err := flavor.List(c, "")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("flavors: %+v\n", flavors)
        for k, v := range flavors.Flavors {
            fmt.Printf("%d: %s\n", k, v.Id)
        }
    return nil
}

func DoListServers(c *client.Client, opts interface{}) error {
    if getField(opts, "Servers").FieldByName("All").Bool() {
        fmt.Printf("all\n")
    }
    servers, err := server.List(c, "")
    fmt.Printf("c: %+v\n\n", *servers)
    return err
}

/*
func DoShowFlavor(c *client.Client, args []string, opts interface{}) error {
    x := getField(opts, "Flavor").FieldByName("F")
    fmt.Printf("DoShowFlavor: %+v\n", x)
    attr := make(flavor.Attr)
    attr["name"] = args[1]
    flavors, err := flavor.Show(c, attr)
    if err != nil {
        fmt.Print("Error: ", err, "\n")
    } else {
        fmt.Printf("flavors: %+v\n\n", flavors.Flavors[0].A)
    }
    return err
}
*/

func DoShowServer(c *client.Client, opts interface{}) error {
    //    if getField(opts, "Servers").FieldByName("All").Bool() {
    //        fmt.Printf("all\n")
    //    }
    attr := make(server.Attr)
    attr["name"] = "npd01"
    servers, err := server.Show(c, attr)
    if err != nil {
        fmt.Print("Error: ", err, "\n")
    } else {
        fmt.Printf("c: %+v\n\n", *servers)
    }
    return err
}

// Display the usage message and exit
func Usage() {
    goptions.PrintHelp()
    os.Exit(1)
}

// Simple helper to set an environment var as default
func SetEnvOpt(env string, opt string) (res string) {
    res = os.Getenv(env)
    if opt != "" {
        res = opt
    }
    return res
}

// Extract auth credentials from command line and environment
func GetAuthCreds(opts OptType) (creds client.Credentials, err error) {
    creds.AuthUrl = SetEnvOpt("OS_AUTH_URL", opts.AuthUrl)
    if creds.AuthUrl == "" {
        err = errors.New("OS_AUTH_URL not found")
    }

    creds.OSAuth.TenantName = SetEnvOpt("OS_TENANT_NAME", opts.Tenant)
    if creds.OSAuth.TenantName == "" {
        err = errors.New("OS_TENANT_NAME not found")
    }

    creds.OSAuth.PasswordCredentials.Username = SetEnvOpt("OS_USERNAME", opts.Username)
    if creds.OSAuth.PasswordCredentials.Username == "" {
        err = errors.New("OS_USERNAME not found")
    }

    creds.OSAuth.PasswordCredentials.Password = SetEnvOpt("OS_PASSWORD", opts.Password)
    if creds.OSAuth.PasswordCredentials.Password == "" {
        // if password == "" here go look in keyring?
    }
    return
}

type OptType struct {
    // TODO(dtroyer): Fix help strings to reflect required args
    Tenant   string `goptions:"--os-tenant, description='Tenant name'"`
    Username string `goptions:"--os-username, description='Username'"`
    Password string `goptions:"--os-password, description='Password'"`
    AuthUrl  string `goptions:"--os-auth-url, description='Authentication URL'"`
    Debug    bool   `goptions:"-x, --debug, description='Enable debugging'"`
    Help     bool   `goptions:"-h, --help, description='Display this help messahe'"`
    Verbose  bool   `goptions:"-v, --verbose, description='Be not quiet with output'"`
    goptions.Verbs
    ListFlavors struct {
        All bool `goptions:"--all, description='List all flavors'"`
    }   `goptions:"listflavors"`
    ListServers struct {
        All bool `goptions:"--all, description='List all servers'"`
    }   `goptions:"listservers"`
    Show struct {
        goptions.Verbs
        Flavor struct {
            F string // `goptions:"--qaz"`
        }   `goptions:"flavor"`
        Server struct {
            S string //`goptions:"-s, --server, description='Server to connect to'"`
        }   `goptions:"server"`
    }   `goptions:"show"`
    Version struct {
        Identity bool `goptions:"--identity, mutexgroup='apiver', description='Show Identity API version'"`
    }   `goptions:"version"`
}

func main() {
    var options OptType = OptType{
        Debug:   false,
        Verbose: false,
    }
    goptions.Parse(&options)

    // Propagate debug setting to packages
    client.Debug = &options.Debug

    if len(os.Args) <= 1 || options.Help {
        Usage()
    }

    // Get our credentials
    creds, err := GetAuthCreds(options)
    if err != nil {
        log.Fatal(err)
    }

    // Make a new client with these creds
    c, err := client.NewClient(creds, nil)
    if err != nil {
        log.Fatal(err)
    }

    if options.Debug {
        fmt.Printf("token: %s\n", c.Token)
        fmt.Printf("servcat: %s\n", c.ServCat)
    }

    fmt.Printf("verbs: %+v\n", options.Verbs)
//    fmt.Printf("obj: %+v\n", options.Show.Verbs)

    switch options.Verbs {
    case "listflavors":
        fmt.Print("Args: ", os.Args[1:], "\n")
        if len(os.Args) > 1 {
            if f, ok := ListCmds["flavors"]; ok {
                err = f(c, &options.ListFlavors)
            } else {
                fmt.Printf("Unknown command: %s %s\n", options.Verbs, os.Args[1])
                fmt.Print("Args: ", os.Args[1:], "\n")
            }
        } else {
            Usage()
        }
        os.Exit(0)
    case "listservers":
        if len(os.Args) > 2 {
            if f, ok := ListCmds["servers"]; ok {
                err = f(c, &options.ListServers)
            } else {
                fmt.Printf("Unknown command: %s %s\n", options.Verbs, os.Args[3])
                fmt.Print("Args: ", os.Args[1:], "\n")
            }
        } else {
            Usage()
        }
        os.Exit(0)
    case "show":
        if len(os.Args) > 2 {
            if f, ok := ShowCmds[string(options.Show.Verbs)]; ok {
                err = f(c, &options.Show)
            } else {
                fmt.Printf("Unknown command: %s %s\n", options.Verbs, os.Args[3])
                fmt.Print("Args: ", os.Args[1:], "\n")
            }
        } else {
            Usage()
        }
        os.Exit(0)
    case "version":
        if options.Version.Identity {
            // Identity API versions
            //                osclib.GetVersions(c.Auth)
            os.Exit(0)
        } else {
            fmt.Printf("Version: %s\n", Version)
            os.Exit(0)
        }
    default:
        fmt.Printf("Unknown verb: %s\n", options.Verbs)
        Usage()
    }

    //    attr := make(server.Attr)
    //    attr["name"] = "npd01"
    //    servers, err := server.Show(c, attr)
    //    if err != nil {
    //        log.Fatal(err)
    //    }
    //    fmt.Printf("c: %+v\n\n", *servers)

}
