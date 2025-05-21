# Use Go base image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy Go module files if using modules
# COPY go.mod go.sum ./
# RUN go mod download

# Copy source code and video file
COPY . .
RUN chmod +x entrypoint.sh

# Build the Go app
# RUN go build -o video-server

# Expose port
EXPOSE 8080

# Run the server
CMD ["/bin/sh", "./entrypoint.sh"]
