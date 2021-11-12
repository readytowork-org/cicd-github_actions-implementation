FROM golang:latest

# Create a buid directory
RUN mkdir /build

# add directory
ADD . /build

# Set the Current Working Directory inside the container
WORKDIR /build

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

EXPOSE 8080

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ./main