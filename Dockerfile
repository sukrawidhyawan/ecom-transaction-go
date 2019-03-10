# transaction Service

# golang builder stage
FROM golang:1.9.0 as builder
RUN mkdir /build_app
WORKDIR /build_app
COPY . .

# TODO: getting the dependencies
RUN env GIT_TERMINAL_PROMPT=1 go get github.com/tripLo-team/ecom-transaction-go github.com/gorilla/mux github.com/jinzhu/gorm
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo main.go

# execution stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /build_app .
RUN ls
EXPOSE 8000
CMD ["./main"]