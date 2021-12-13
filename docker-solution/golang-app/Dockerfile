# builder image
FROM golang:1.17.5-alpine3.15 as builder
RUN mkdir /build
ADD *.go /build/
ADD *.mod /build/
WORKDIR /build
RUN go get -u -v -f all && CGO_ENABLED=0 GOOS=linux go build -a -o golang-app .

# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/golang-app .

EXPOSE 8005/tcp

# executable
ENTRYPOINT [ "./golang-app" ]
