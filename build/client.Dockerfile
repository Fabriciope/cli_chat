FROM golang:1.22.1

WORKDIR /app

COPY ./../client/ /app/client
COPY ./../cmd/client/ /app/cmd/
COPY ./../pkg/ /app/pkg/
COPY ./../go.mod ./../go.sum /app/

RUN go mod download && go mod verify
RUN go build -o /go/bin/app /app/cmd/main.go

CMD ["/go/bin/app"]
