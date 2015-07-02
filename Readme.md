# Go Insteon!
## Insteon library written poorly in Go

go get github.com/swisskid/go-insteon

then import "github.com/swisskid/go-insteon/insteon"

To make this work, you need to create secrets.go with the following:
package insteon

var Access_Token string  = "<access_token>"
var Refresh_Token string  = "<refresh_token>"
var Client_Id string  = "<api_key>"
var Insteon_Url string  = "https://connect.insteon.com/api/v2/"

Now, library doesn't currently know how to refresh a token... so you're kinda SOL there.
I'll eventually write this so it reads it out of a file, and learns to store/refresh the access token.

Maybe use
https://github.com/google/go-querystring


