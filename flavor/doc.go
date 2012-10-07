// flavor/doc.go - roscoe flavor package documentation

/*

'flavor' package implements the flavor OpenStack Compute APIs

Currently supported API version(s): v2

Example:

    var creds osclib.Creds
    c, err := client.NewClient(creds)
    flavors, err := flavor.List(c, "")

*/
package flavor
