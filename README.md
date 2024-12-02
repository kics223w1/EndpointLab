# EndpointLab

An alternative to [httpbin.org](https://github.com/postmanlabs/httpbin) in Gin.

## Quick Start

```bash
docker pull viethuy/endpointlab
docker run -p 8080:8080 viethuy/endpointlab
```

## Swagger UI

Start the UI:

```bash
make run
```

Access it at: `http://localhost:8080/swagger/index.html`

Update `swagger.json`:

```bash
make swagger
```

## Build and Distribute

```bash
make dockerDistribute
```
