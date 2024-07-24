# Gateway

Minimal API gateway built on top of Fasthttp for proxying requests to [@abyan-dev](https://github.com/abyan-dev) microservices.

To run the app, use:

```
go run ./cmd/api
```

The gateway leverages Docker's DNS to proxy requests to corresponding services by their container names. It is also equipped with middlewares for authorization, TLS, and HTTP logger.