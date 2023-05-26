FROM golang:1.18 as builder

ENV GOPATH=""
ENV GOMODULE="on"

COPY . .

RUN go mod download
RUN make build

FROM bitnami/minideb
COPY --from=builder /go/build/simd /bin/simd

ENTRYPOINT ["simd"]

