## Gin - Example correlation id middleware usage 
creates a gin router, listening on port 300,  sets up the correlation id middleware and registers the handlers

### Build the server
From the project root

    go build cmd/examples/gin/gin.go
<br />

---

### Run the server
    ./gin

<br />

---


### Curl, containing `NO` correlation id in the request 
Curl the server, with headers flag enabled, with no correlation id, one will be created and added to the response

    curl -i localhost:8000/hello

The server returns the correlation id header (in the example, named `my-trace-header`, but this is configurable) 
and value in the response headers (and, in this case, in the body)

```http response
HTTP/1.1 200 OK
My-Trace-Header: b63a65cc-20fa-4b17-97ad-b796bdb6d338      <--- correlation id header
Date: Thu, 01 Jun 2023 01:13:22 GMT
Content-Length: 107
Content-Type: text/plain; charset=utf-8

Hello World!
correlation id header in  = 
correlation id header out = b63a65cc-20fa-4b17-97ad-b796bdb6d338
```

<br />

---

### Curl, with a correlation id in the request
Curl the server, adding in the correlation id header

    curl -i localhost:8000/hello -H "my-trace-header: hello-world-1"

The server will return the correlation id header and value in the response headers (and, in this case, in the body)

```http response
HTTP/1.1 200 OK
My-Trace-Header: hello-world-1      <--- correlation id header
Date: Thu, 01 Jun 2023 01:14:38 GMT
Content-Length: 97
Content-Type: text/plain; charset=utf-8

Hello World!
correlation id header in  = hello-world-1
correlation id header out = hello-world-1
```
