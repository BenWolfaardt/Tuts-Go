# Protbuffer

Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

* [Github - Main Project](https://github.com/protocolbuffers/protobuf)
* [GitHub - Go Project](https://github.com/protocolbuffers/protobuf-go)
* [Language Guide](https://developers.google.com/protocol-buffers/docs/overview)
* [Go support for Protocol Buffers](https://pkg.go.dev/google.golang.org/protobuf#section-readme)

# gRPC 

gRPC is a modern open source high performance **Remote Procedure Call (RPC)** framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

[This page](https://grpc.io/docs/what-is-grpc/introduction/) introduces you to gRPC and protocol buffers. gRPC can use protocol buffers as both its Interface Definition Language (IDL) and as its underlying message interchange format.

* [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
* [gRPC - Documentation](https://grpc.io/docs/)
* [gRPC - Go Documentation](https://grpc.io/docs/languages/go/)

# Working With Protobuf

Using [Protocol Buffer Basics tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)

Build the Go samples with "make go".  This creates the following
executable files in the current directory:
```s
    add_person_go      list_people_go
```

To run the example:
```s
    ./add_person_go addressbook.data
```

to add a person to the protocol buffer encoded file addressbook.data.  The file
is created if it does not exist.  To view the data, run:
```s
    ./list_people_go addressbook.data
```

# Install & Setup

1. Follow instructions in the [README.md](https://github.com/protocolbuffers/protobuf/blob/master/README.md) (from the `protobuf` GitHub [repo](https://github.com/protocolbuffers/protobuf/))to install protoc. Then
   1. Download the Ubuntu x86-64 version from: [releases](https://github.com/protocolbuffers/protobuf/releases)
   2. Extract and paste in `usr/local/bin` and usr`/local/include`
2. install the Go protoc plugin (protoc-gen-go)
   * `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
