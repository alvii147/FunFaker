FROM golang:1.19 as go-builder
# Create and change to the app directory.
WORKDIR /go/src/app
# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download
# Copy local code to the container image.
COPY . ./
# Build the binary.
RUN go build -o main main.go

FROM ubuntu:latest
# Create and change to the app directory.
WORKDIR /go/src/app
# Copy the binary to the production image from the builder stage.
COPY --from=go-builder /go/src/app/ ./
# Run the web service on container startup.
CMD ["./main", "server"]
