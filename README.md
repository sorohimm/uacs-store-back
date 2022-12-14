# E-commerce backend template with JWT Auth

[![Go Report Card](https://goreportcard.com/badge/github.com/sorohimm/uacs-store-back)](https://goreportcard.com/report/github.com/sorohimm/uacs-store-back)

Backend part for an online store written in Golang. Supports HTTP and RPC requests.

[API-Reference (Swagger)](https://app.swaggerhub.com/apis/sorohimm3/UACS/v1)

### Dependencies

- PostgreSQL

### Local development environment

deploying:
```
$ make compose-up 
```

stopping:
```
$ make compose-down
```

migrate (required for the first time):
```
$ make migrate-up
```

### Local run

```
$ make
```
This command will assemble the auth and store services

Then run our services like this
```
$ ./build/uacs-store --pg.schema=store
```
```
$ ./build/uacs-auth --pg.schema=users --jwt.secret=<secret>
```
