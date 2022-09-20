# twotop-backend

### Running locally

With Docker, spin up the db as a container and spin up the service as a process

* [Install Podman](https://github.com/containers/podman-compose#installation)

* `podman-compose up`

* `export INTERNAL_URL="postgres://twotop:twotop@localhost:5432/twotop"`

* `go run main.go`
