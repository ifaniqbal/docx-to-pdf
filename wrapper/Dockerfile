# Use the official Golang image as the parent image
FROM golang:1.20

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 80

# Run the application
CMD ["./main"]