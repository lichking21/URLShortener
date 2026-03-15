# Go URL Shortener

A high-performance, lightweight URL shortener service (similar to bit.ly) built with Go. 

## Features

* **RESTful API**: Simple and intuitive endpoints for creating and redirecting URLs.
* **Base62 Encoding**: Generates short, collision-free URL codes based on database sequences.
* **High Performance**: Uses **Redis** as an in-memory cache (Cache-Aside pattern) to serve frequently accessed links in milliseconds.
* **Persistent Storage**: Uses **PostgreSQL** for reliable data storage.
* **Containerized**: Fully dockerized environment using `podman-compose` / `docker-compose` for easy deployment.

## Tech Stack

* **Language**: Go (1.26+)
* **Database**: PostgreSQL 16 (using `pgx` / `sqlx`)
* **Cache**: Redis 7 (using `go-redis`)
* **Infrastructure**: Podman / Docker Compose
* **Routing**: Standard Go `net/http` ServeMux (Go 1.22+ features)

## Quick Start
`curl -X POST http://localhost:8080/api/shorten \
     -H "Content-Type: application/json" \
     -d '{"url": "https://go.dev/doc/"}'`

### Prerequisites
* [Go](https://golang.org/doc/install) (if running locally without containers)
* [Podman](https://podman.io/) or [Docker](https://www.docker.com/) with Compose
