# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.12 as builder

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# RUN go test -v ./services/...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrate ./migration


######## Start a new stage from scratch #######
FROM alpine:latest  

WORKDIR /root/

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add  --no-cache ffmpeg

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"] 