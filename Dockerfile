# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only necessary files to the container
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the local code to the container
COPY . .

# Build the Go application and place the executable in /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Move the executable and .env outside the working directory
RUN mv main /

# Delete the working directory
RUN rm -rf /app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/main"]
