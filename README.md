roscoe: OpenStack Client
========================

Description
-----------

roscoe is an OpenStack client written in Go.  It is a fresh implementation
of the OpenStack client libraries and a single command-line binary.

Usage
-----

It's so early that there isn't any yet.  Look in ``examples`` for now.

Acquiring
---------

    sudo apt-get install golang
    export GOPATH=$HOME/go    # or whatever
    mkdir -p $GOPATH/src
    cd $GOPATH/src
    git clone https://github.com/asdfio/roscoe.git
    cd roscoe

Building
--------

    go build
    ./roscoe

Using
-----

Really-low-level:

    var creds osclib.Creds
    c, err := client.NewClient(creds)
    servers, err := server.List(c, "")

Error handling omitted.

Examples
--------

token
_____

Simple example to retrieve a token from an OpenStack Identity service
using the OpenStack auth environment variables (OS_TENANT_NAME, OS_USERNAME,
OS_PASSWORD, OS_AUTH_URL).

    go build examples/token
    ./token

Dependencies
------------

roscoe uses goopt "github.com/droundy/goopt" for command-line parsing
