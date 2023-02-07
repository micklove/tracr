# tracr
simple lib for adding a correlation id / trace id to the context 

<img src='./gopher.png' width='300'>

---

# correlation-id
tracr is a Go lib that allows clients to add / retrieve a correlation id header (with header name of their choice) to / from the current go context.

The library provides the following:

---

## 1. Middleware for incoming requests
Middleware method will look for a correlation id header in the current request.
If no correlation id is found, a new correlation id will be generated and added to the request context.

---

#### Correlation ID generation
nb: If no correlation ID is found in the context, a new correlation id will be generated and added.
The library, https://github.com/gofrs/uuid, will be used to generate the correlation id.
In future releases, we may consider adding a method to allow tracr users to provide their own method for generating correlation ids.

<br />

### Preferred Header Name
During middleware instantiation, users can provide the preferred name for the correlation id header.
e.g. `x-correlation-id, x-request-id, trace-id, trace_id`, etc...

If no names are provided, the default header will be used: (nb: header names are case insensitive)

    x-correlation-id


---


## 2. Adding correlation id to outgoing requests
See request.go. A method is provided to allow the addition of a correlation id to outgoing requests

---

<br />

## Installation
Installation with go get.

    go get -u github.com/micklove/tracr

---

<br />

### TODO Examples

Add examples

---

### TODO - Diagram
add mermaid diagram(s) 

---

## References
Pages reviewed / consulted / borrowed from, when coming up with this lib

TODO - clean up

* https://github.com/auroratechnologies/vangoh
* https://kevin.burke.dev/kevin/how-to-write-go-middleware/
* https://community.developers.refinitiv.com/questions/57029/how-to-generate-signature-in-golang-any-one-help-m.html
* https://www.reddit.com/r/golang/comments/mnht8z/getting_response_headers_and_body_in_middleware/
* https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
* https://upgear.io/blog/golang-tip-wrapping-http-response-writer-for-middleware/
* https://github.com/alexsergivan/blog-examples/blob/master/middleware/main.go

### Issues to consider


### TODO - contributing
