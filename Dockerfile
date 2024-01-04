FROM golang::1.21.5-alpine
COPY . /go/src/github.com/RyoMasumura1201/go-todo

WORKDIR /go/src/github.com/RyoMasumura1201/go-todo

RUN go mod tidy
RUN go install github.com/RyoMasumura1201/go-todo

ENTRYPOINT /go/bin/go-todo

# Document that the service listens on port 8080.
EXPOSE 8080