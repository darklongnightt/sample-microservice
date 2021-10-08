# sample-microservice

Sample microservice written in golang and docker with the following examples

1. Restful APIs

```
/homepage

- Endpoint that returns a simple json
- Endpoint to handle file upload
```

```
/productpage

- Endpoint to write a product to db
- Endpoint to read a product from db
```

Folder structures by logical modules for each group of REST endpoints
Example:

```
    /homepage
        - handlers
        - middlewares
        - funcs
        - models
```

2. TLS (HTTPS) Server

   /server contains TLS configs

   more info on setting this configs: https://blog.cloudflare.com/exposing-go-on-the-internet/

   /certs contains self signed certificate and private key to stimulate HTTPS connections

   **Note:** Not to be used in live server

3. Database

   /db folder contains all sql db models such as products
