// servers/doc.go - roscoe server package documentation

/*

'server' package implements the server OpenStack Compute APIs

Currently supported API version(s): v2

Example:

    var creds osclib.Creds
    c, err := client.NewClient(creds)
    servers, err := server.List(c, "")

*/
package server
