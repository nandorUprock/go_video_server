# Use Go base image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy source code
COPY entrypoint.sh .
COPY main.go .
RUN chmod +x entrypoint.sh

# Expose port
EXPOSE 8080

# Run the server
CMD ["/bin/sh", "./entrypoint.sh"]
