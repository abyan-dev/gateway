# Gateway

![Tests](https://github.com/abyan-dev/gateway/actions/workflows/ci.yaml/badge.svg) [![codecov](https://codecov.io/gh/abyan-dev/gateway/graph/badge.svg?token=wQp0whBQ6W)](https://codecov.io/gh/abyan-dev/gateway) [![Go Report](https://goreportcard.com/badge/abyan-dev/gateway)](https://goreportcard.com/report/abyan-dev/gateway) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/abyan-dev/gateway/blob/main/LICENSE)

Minimal API gateway built on top of Fasthttp for proxying requests to [@abyan-dev](https://github.com/abyan-dev) microservices.

To run the app, use:

```
go run ./cmd/api
```

The gateway leverages Docker's DNS to proxy requests to corresponding services by their container names. It is also equipped with middlewares for authorization, TLS, and HTTP logger.

## Integration

This gateway uses Fiber's `Proxy` middleware, so all you need to do is catch all routes after your service's domain using a wildcard:

```go
app.All("<url_in_gateway>", func(c *fiber.Ctx) error {
    targetURL := "<target_url>" + c.OriginalURL()[len("<url_in_gateway>"):]
    if err := proxy.Do(c, targetURL); err != nil {
      return err
    }
    c.Response().Header.Del(fiber.HeaderServer)
    return nil
})
```
