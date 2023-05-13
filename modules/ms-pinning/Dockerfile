# Use an official Go runtime as a parent image
FROM golang:alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Compile the Go program
RUN go build -o hello .

# Run the command by default when the container starts
CMD ["./hello"]
