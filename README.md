# go-project-demo

This project is for demonstrating go project purpose

This repo is one of the microservices. Other related repo are <b>[elasticsearch-monstache-demo](https://github.com/alanyeung95/elasticsearch-monstache-demo)</b> and <b>[vue-project-demo](https://github.com/alanyeung95/vue-project-demo)</b>

Please setup the project under this order:

1. go-project-demo
2. elasticsearch-monstache-demo
3. vue-project-demo

## Quick usage

### docker compose

Run the API server

```
docker-compose up
```

Run unit test

```
make cover
```

## Todo list:

- [x] Create item API
- [x] Get item API
- [x] Provide swagger
- [x] Add sentry to monitor error (may be)
- [ ] MongoDB indexing
- [x] Gomock test
- [x] User auth with JWT

## Useful information

### Gomock

1.  Gomock tutorial: https://blog.codecentric.de/en/2017/08/gomock-tutorial/

### Ginkgo

A Golang Behavior-Driven Development ("BDD") testing framework

Example

```
Describe("the strings package", func() {
  Context("strings.Contains()", func() {
    When("the string contains the substring in the middle", func() {
      It("returns `true`", func() {
        Expect(strings.Contains("Ginkgo is awesome", "is")).To(BeTrue())
      })
    })
  })
})
```

### Gomega

Gomega is a go matcher used in go test asserting

Example

```
Equal(...)
BeEquivalentTo(...)
BeNil()
BeZero()
BeTrue()
BeFalse()
```

## Notes

### gRPC

In gRPC with Go, when you use the Protocol Buffers compiler (`protoc`) to generate Go code from `.proto` files, you end up with two types of files:

1. **`helloworld.pb.go`**: This file contains:

   - Struct definitions for each of your Protocol Buffers messages.
   - Code to encode (marshal) and decode (unmarshal) these messages to and from binary format.
   - Any enum definitions from your `.proto` file as Go constants.
   - file_helloworld_proto_rawDesc is the whole `.proto` file in binary format

2. **`helloworld_grpc.pb.go`**: This file is specifically for gRPC and includes:
   - Client and server interfaces for the gRPC services defined in your `.proto` file.
   - Code that helps implement these services on the server-side and call them on the client-side.

Essentially, `xxx.pb.go` deals with data structures, while `xxx_grpc.pb.go` handles the communication aspects using those structures in gRPC.

## Troubleshooting

### DeadlineExceeded

```
 2024/08/18 05:35:38 Error greeting aac: rpc error: code = DeadlineExceeded desc = context deadline exceeded
```

This error message indicates that the operation exceeded the time allowed by the deadline set in the context. This is commonly encountered in network and RPC operations where you've defined a maximum duration that an operation can take, and the operation did not complete within this timeframe.

Common Causes

1. Network Delays: The server might be taking too long to respond due to network latency, or the server itself might be overloaded and thus slow to handle requests.

2. Server Processing Time: The server-side operation might be more time-consuming than anticipated. This could be due to data processing requirements, database access delays, or other computational tasks.

3. Improper Timeout Settings: The client might have set a deadline that is too short for the operation to realistically complete in normal conditions.
