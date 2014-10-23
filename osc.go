// osc.go - roscoe command line

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/voxelbrain/goptions"

	"roscoe/auth"
	// "roscoe/flavor"
	"roscoe/server"
	"roscoe/session"
)

const (
	Version = "0.2"
)

type CmdFunc func(c *session.Session, a *auth.Auth, opts interface{}) error

var FlavorCmds map[string]CmdFunc = map[string]CmdFunc{
// "list": DoFlavorList,
// "show": DoFlavorShow,
}

var ServerCmds map[string]CmdFunc = map[string]CmdFunc{
	"list": DoServerList,
	"show": DoServerShow,
}

func getField(opts interface{}, name string) (r reflect.Value) {
	v := reflect.ValueOf(opts).Elem()
	if v.Kind() == reflect.Struct {
		r = v.FieldByName(name)
		return r
	}
	return r
}

/*
func DoFlavorList(c *session.Session, opts interface{}) error {
    if getField(opts, "All").Bool() {
        fmt.Printf("all\n")
    }
    flavors, err := flavor.List(c, "")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("flavors: %+v\n", reflect.TypeOf(flavors))
        for _, l := range *flavors {
            for k, v := range l {
                fmt.Printf("f: %+v\n", v[k])
            }
        }
    return nil
}
*/

/*
func DoFlavorShow(c *session.Session, args []string, opts interface{}) error {
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

func DoServerList(s *session.Session, a *auth.Auth, opts interface{}) error {
	if getField(opts, "Servers").FieldByName("All").Bool() {
		fmt.Printf("all\n")
	}
	servers, err := server.List(s, "")
	fmt.Printf("s: %+v\n\n", *servers)
	return err
}

func DoServerShow(s *session.Session, a *auth.Auth, opts interface{}) error {
	name := getField(opts, "Show").FieldByName("S")
	fmt.Printf("server: %s\n", name)
	attr := make(server.Attr)
	attr["name"] = "x"
	servers, err := server.Show(s, a, attr)
	if err != nil {
		fmt.Print("Error: ", err, "\n")
	} else {
		fmt.Printf("s: %+v\n\n", *servers)
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
func GetAuthCreds(opts OptType) (creds auth.UserPassV2, err error) {
	creds.AuthUrl = SetEnvOpt("OS_AUTH_URL", opts.AuthUrl)
	if creds.AuthUrl == "" {
		err = errors.New("OS_AUTH_URL not found")
	}

	creds.OSAuth.ProjectName = SetEnvOpt("OS_PROJECT", opts.Project)
	if creds.OSAuth.ProjectName == "" {
		creds.OSAuth.ProjectName = SetEnvOpt("OS_PROJECT_NAME", opts.Project)
		if creds.OSAuth.ProjectName == "" {
			err = errors.New("OS_PROJECT not found")
		}
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
	AuthUrl  string `goptions:"--os-auth-url, description='Authentication URL'"`
	Project  string `goptions:"--os-project, description='Project name'"`
	Username string `goptions:"--os-username, description='Username'"`
	Password string `goptions:"--os-password, description='Password'"`
	Debug    bool   `goptions:"-x, --debug, description='Enable debugging'"`
	Help     bool   `goptions:"-h, --help, description='Display this help messahe'"`
	Verbose  bool   `goptions:"-v, --verbose, description='Be not quiet with output'"`

	goptions.Verbs

	Flavor struct {
		goptions.Verbs
		List struct {
			goptions.Remainder
			All bool `goptions:"--all, description='List all flavors'"`
		} `goptions:"list"`
		Show struct {
			goptions.Remainder
			F string // `goptions:"--qaz"`
		} `goptions:"show"`
	} `goptions:"flavor"`

	Server struct {
		goptions.Verbs
		List struct {
			goptions.Remainder
			All bool `goptions:"--all, description='List all servers'"`
		} `goptions:"list"`
		Show struct {
			// goptions.Remainder
			All bool   `goptions:"--all, description='List all servers'"`
			S   string //`goptions:"-s, --server, description='Server to connect to'"`
		} `goptions:"show"`
	} `goptions:"server"`

	Version struct {
		Identity bool `goptions:"--identity, mutexgroup='apiver', description='Show Identity API version'"`
	} `goptions:"version"`
}

func main() {
	var options OptType = OptType{
		Debug:   false,
		Verbose: false,
	}
	goptions.Parse(&options)

	// Propagate debug setting to packages
	session.Debug = &options.Debug

	if len(os.Args) <= 1 || options.Help {
		Usage()
	}

	// Get our credentials
	creds, err := GetAuthCreds(options)
	if err != nil {
		log.Fatal(err)
	}

	auth_ref, err := creds.GetAuthRef()

	// Make a new client with these creds
	c, err := session.NewSession(creds.AuthUserPassV2, nil)
	if err != nil {
		log.Fatal(err)
	}

	if options.Debug {
		fmt.Printf("token: %s\n", c.Token)
	}

	// Option debugging
	fmt.Printf("object: %+v\n", options.Verbs)
	fmt.Printf("verb: %+v\n", options.Server.Verbs)
	fmt.Printf("args: %+v\n", options.Server.Show)

	switch options.Verbs {
	case "flavor":
		if len(os.Args) > 2 {
			if f, ok := FlavorCmds[string(options.Flavor.Verbs)]; ok {
				err = f(c, auth_ref, &options.Flavor)
			} else {
				fmt.Printf("Unknown command: %s %s\n", options.Verbs, os.Args[3])
				fmt.Print("Args: ", os.Args[1:], "\n")
			}
		} else {
			Usage()
		}
		os.Exit(0)
	case "server":
		if len(os.Args) > 2 {
			if f, ok := ServerCmds[string(options.Server.Verbs)]; ok {
				err = f(c, auth_ref, &options.Server)
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
