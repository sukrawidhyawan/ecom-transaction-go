# Products Service

# golang builder stage
FROM golang:1.9.0 as builder
COPY . .

# TODO: getting the dependencies
RUN env GIT_TERMINAL_PROMPT=1 go get github.com/tripLo-team/ecom-transaction-go github.com/gorilla/mux github.com/jinzhu/gorm
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# execution stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/tripLo-team/ecom-transaction-go .
EXPOSE 8000
CMD ["./main"]