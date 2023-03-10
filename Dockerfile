FROM alpine:3.16 as root-certs
#RUN apk --no-cache add ca-certificates
#RUN addgroup -g 1001 app
#RUN adduser app -u 1001 -D -G app /home/app

# # specify the base image for GO app
FROM golang:1.17 as builder
# # Create /app dir within the image to hold our application source code.
WORKDIR /home/app
# # copy root certificate from previous build 
#COPY --from=root-certs  /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
# # copy GO mod into WORKDIR
COPY go.mod .
# # install dependencies
RUN go mod download
# # copy go files into WORKDIR
COPY main.go .
# # build the app within optional configs
#RUN go build -o /go-test-app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-test-app
# # tell docker to listen on the port at runtime
# # command to be used to execute when image is used to start the container.
ENTRYPOINT ["/go-test-app"]
