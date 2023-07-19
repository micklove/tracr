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
X-Correlation-Id: 68a914cd-591c-4701-b208-98c2181d8a2e
Date: Wed, 19 Jul 2023 13:49:21 GMT
Content-Length: 41
Content-Type: text/plain; charset=utf-8

cid: 68a914cd-591c-4701-b208-98c2181d8a2e
```

<br />

---

### Curl, with a correlation id in the request
Curl the server, adding in the correlation id header

    curl -i localhost:8087/hello -H "x-correlation-id: hello-world-1"

The server will return the correlation id header and value in the response headers (and, in this case, in the body)

```http response
HTTP/1.1 200 OK
X-Correlation-Id: hello-world-1
Date: Wed, 19 Jul 2023 13:48:57 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

cid: hello-world-1

```
