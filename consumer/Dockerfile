# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.17 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY ./events ../events

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./consumer/go.mod ./

RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY ./consumer .

# RUN swag init

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Add files to the image
ADD ./consumer/config config

# Command to run the executable
CMD ["./main"]

