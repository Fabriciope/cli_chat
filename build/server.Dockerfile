FROM golang:1.22.1

WORKDIR /app

COPY ./server/ /app/
COPY ./shared/ /app/shared/
COPY ./go.mod /app/

RUN go mod tidy && go mod download
RUN go build -o /app server

CMD ["./server"]
