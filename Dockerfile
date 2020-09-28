# Build:    $ docker build -t my-go-app . 
# Run:      $ docker run -it -p 8080:8080 my-go-app

FROM golang:1.15.2-alpine AS build-env
ENV GO111MODULE=on

# Allow Go to retrive the dependencies for the build step without caching
RUN apk add --no-cache git

# Secure against running as root
RUN adduser -D -u 10000 xavier
RUN mkdir /go/microservice && chown xavier /go/microservice
USER xavier

# Create work directory and download dependencies before mirroring
WORKDIR /go/microservice
COPY . .
RUN go mod download

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o main .

# Create env used in application and expose port used
ENV LOCALHOST_CERT "./certs/localhost.crt"
ENV LOCALHOST_PRIVATE_KEY "./certs/localhost.key"
EXPOSE 8080
CMD ["/go/microservice/main"]

