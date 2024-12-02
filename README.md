# EndpointLab

An alternative to [httpbin.org](https://github.com/postmanlabs/httpbin) in Gin.

## Quick Start

```bash
docker pull viethuy/endpointlab
docker run -p 8080:8080 viethuy/endpointlab
```

## Swagger UI

Access the Swagger UI:

- **Online**: Visit [SwaggerHub](<https://app.swaggerhub.com/apis-docs/HuyCao(John)/EndpointLab/1.0>).

- **Locally**:

  1. Start the UI:
     ```bash
     make run
     ```
  2. Open in browser: `http://localhost:8080/swagger/index.html`

- **Update `swagger.json`**:
  ```bash
  make swagger
  ```

## Build and Distribute

```bash
make dockerDistribute
```
