## Stdlib- Example correlation id middleware usage 
creates a standard lib http server, listening on port 8087,  sets up the correlation id middleware and registers the handlers

### Build the server
From the project root

    go build cmd/examples/stdlib/stdlib.go
<br />

---

### Run the server
    ./stdlib

<br />

---


### Curl, containing `NO` correlation id in the request 
Curl the server, with headers flag enabled, with no correlation id, one will be created and added to the response

    curl -i localhost:8087

The server returns the correlation id header (in the example, named `my-trace-header`, but this is configurable) 
and value in the response headers (and, in this case, in the body)

```http response
HTTP/1.1 200 OK
My-Trace-Header: 0c7928a0-4fb5-425a-98b2-019fb0317ee9
Date: Tue, 18 Jul 2023 12:50:24 GMT
Content-Length: 41
Content-Type: text/plain; charset=utf-8

cid: 0c7928a0-4fb5-425a-98b2-019fb0317ee9

```

<br />

---

### Curl, with a correlation id in the request
Curl the server, adding in the correlation id header

    curl -i localhost:8087/hello -H "my-trace-header: hello-world-1"

The server will return the correlation id header and value in the response headers (and, in this case, in the body)

```http response
HTTP/1.1 200 OK
My-Trace-Header: hello-world-1
Date: Tue, 18 Jul 2023 12:49:40 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

cid: hello-world-1
```
