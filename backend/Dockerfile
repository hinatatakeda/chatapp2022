# FROM golang:1.17.7-alpine
FROM golang:1.17.7-alpine
RUN apk update && apk add git
WORKDIR /go/src
ADD . /go/src/



CMD ["go", "run", "main.go"]



# FROM golang:1.18-alpine
# FROM alpine:latest


# WORKDIR /app/go/base

# COPY go.mod .
# COPY go.sum .

# EXPOSE 8080
# ENV GO111MODULE=on

# RUN go mod download
# RUN go install github.com/cosmtrek/air@latest
# COPY . .





# CMD ["go","run","./main.go"]