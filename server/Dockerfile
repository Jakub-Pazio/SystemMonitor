# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 1323 to the outside world
EXPOSE 1323

# Run the executable
CMD ["./main"]

