roscoe: OpenStack Client
========================

Description
-----------

roscoe is an OpenStack client written in Go.  It is a fresh implementation
of the OpenStack client libraries and a single command-line binary.

Usage
-----

It's so early that there isn't any yet.  ``./osc -h`` should be helpful, or
look in ``examples``.

Acquiring
---------

Set up Go workspace and get roscoe::

    sudo apt-get install golang
    export GOPATH=$HOME/go    # or whatever
    mkdir -p $GOPATH/src
    cd $GOPATH/src
    git clone https://github.com/asdfio/roscoe.git
    cd roscoe

Get dependencies::

    go get github.com/voxelbrain/goptions
    go install github.com/voxelbrain/goptions

Building
--------

    go build
    ./osc

Using
-----

Library API (error handling omitted)::

    var creds osclib.Creds
    c, err := client.NewClient(creds)
    servers, err := server.List(c, "")

Command-line:

    osc list servers

Examples
--------

token
~~~~~

Simple example to retrieve a token from an OpenStack Identity service
using the OpenStack auth environment variables (OS_TENANT_NAME, OS_USERNAME,
OS_PASSWORD, OS_AUTH_URL).

    go build examples/token
    ./token

mystuff
~~~~~~~

Quick list of servers, flavors, images, etc from OpenStack.

Dependencies
------------

roscoe uses github.com/voxelbrain/goptions for command-line parsing
