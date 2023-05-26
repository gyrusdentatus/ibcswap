FROM golang:1.20.4-bullseye AS builder

ENV GOPATH=""  
ENV GOMODULE="on"  

COPY go.mod go.sum ./

RUN go mod download  

COPY testing testing  
COPY modules modules  

COPY Makefile .  

RUN make build  

FROM ubuntu:20.04

# Install necessary packages and clean up
RUN apt-get update && apt-get install -y \
    && rm -rf /var/lib/apt/lists/* 

COPY --from=builder /go/build/simd /bin/simd  

ENTRYPOINT ["simd"]
