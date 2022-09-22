# twotop-backend

### Running locally

With Docker, spin up the db as a container and spin up the service as a process

* [Install Podman](https://github.com/containers/podman-compose#installation)

* `podman-compose up`

* `export INTERNAL_URL="postgres://twotop:twotop@localhost:5432/twotop"`

* `go run main.go`


### Running locally for MacOS (M1)

With Docker, spin up the db as a container and spin up the service as a process

* Install podman with brew: `brew install podman`

* [Install podman-compose](https://github.com/containers/podman-compose#installation)

* `podman machine start`

* `podman-compose up`

* `export INTERNAL_URL="postgres://twotop:twotop@localhost:5432/twotop"`

* `go run main.go`
