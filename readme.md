# Go tips

To initialise a project outside the Go path use `go mod init <module_name>`

# To run program

/Users/Michael.Caley/go/bin/CompileDaemon -command="./go-hue-controls"

Starts an api listen on port 3000 (referenced in .env) http://localhost:3000

# Hue developer guide

Getting started: https://developers.meethue.com/develop/get-started-2

# Note

{"devicetype":"my_hue_app#mac-michael"}

# How to use

Send a get request to host url with path `/light` and light no.

Full example: `localhost:3000/light/5`

Each request to this endpoint will toggle the light between active states.