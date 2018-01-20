# Networking deepdive—From net.Dial to gRPC

FOSDEM 2018 [Go devroom](https://fosdem.org/2018/schedule/event/networking/).

Many code examples based on [Network Programming with Go by Jan Newmarch](http://tumregels.github.io/Network-Programming-with-Go/)

## Part 1: Basics

We start off with a review of the stdlib net package and its sub-packages, walk through common use cases, patterns and challenges:

- Discussion of net and sub-packages (main types, interfaces, understanding what is available)
- Socket-level programming
- Timeouts, retries
- Security
- DNS

## Part 2: Best practices and beyond stdlib

In the second part we focus on best practices using the stdlib and community-provided frameworks for networking:

- Best practices writing networked applications
- Review of non-stdlib packages useful for networking: golang.org/x/net/websocket, gorilla/websocket, gRPC
