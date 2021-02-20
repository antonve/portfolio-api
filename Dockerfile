FROM golang:1.16 as build
WORKDIR /base
COPY . .
RUN go mod download
RUN go install -v ./...

# Create production container
FROM alpine:3.7
COPY --from=build /go/bin/server /go/bin/migrate /usr/bin/

# Run the app
ENTRYPOINT ["server"]
