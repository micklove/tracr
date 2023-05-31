# tracr
simple lib for retrieving (or creating) and adding a correlation id / trace id to the context 

<img src='./gopher.png' width='300'>

---

TODO - Build / Release Pipe
---

### tl;dr ?
Have a look at the examples in [cmd/examples/chi](cmd/examples/chi.go)

# correlation-id
tracr is a Go lib that allows clients to add / retrieve a correlation id header (with header name of their choice) to / from the current go context.

The library provides the following:

---


## 1. Middleware for incoming requests
The middleware checks for a correlation ID header in the current request. If none is found, it generates a new correlation ID and adds it to the request context.

---

#### Correlation ID generation
nb: If no correlation ID is found in the context, a new correlation id will be generated and added to the context:
1. Users can provide their own correlation id generator function, when instantiating the middleware.
OR
2. The library, https://github.com/gofrs/uuid, will be used by default to generate the correlation id (in the form of a uuid)

<br />

### Preferred Header Name
During middleware instantiation, users can provide a func that will return the preferred name for the correlation id http header.
e.g. `x-correlation-id, x-request-id, trace-id, trace_id`, etc...

If no names are provided, the default header will be used: (nb: header names are case insensitive)

    x-correlation-id


---


## 2. Adding correlation id to outgoing requests
See [request.go](/internal/tracr/request.go). 

A method is provided to allow the addition of a correlation id (and correlation id http header) to outgoing requests

---

<br />

## Installation
Installation with go get.

    go get -u github.com/micklove/tracr

---

<br />

### Examples

See [the examples folder](/cmd/examples/)

---

### TODO - Diagram
add mermaid diagram(s) 

---

## References
Pages reviewed / consulted / borrowed from, when coming up with this lib

* https://github.com/auroratechnologies/vangoh
* https://kevin.burke.dev/kevin/how-to-write-go-middleware/
* https://community.developers.refinitiv.com/questions/57029/how-to-generate-signature-in-golang-any-one-help-m.html
* https://www.reddit.com/r/golang/comments/mnht8z/getting_response_headers_and_body_in_middleware/
* https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
* https://upgear.io/blog/golang-tip-wrapping-http-response-writer-for-middleware/
* https://github.com/alexsergivan/blog-examples/blob/master/middleware/main.go

### Issues to consider
If an error occurs in the middleware func, the current implementation will just write an http 500 to to the response, if there is an issue retrieving the correlation id - TODO - add a method to override this behaviour

### TODO - contributing
