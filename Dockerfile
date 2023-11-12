#Set the base image
FROM golang:latest

# Set metadata labels
LABEL maintainer = "Rami, Hasan, Hussain"
LABEL description = "ASCII art web Docker image"

# Set the working directory
WORKDIR /ascii-art-web-dockerize

#Copy the application code
COPY . .

WORKDIR /ascii-art-web-dockerize/cmd/asciiartweb

# Build the application
RUN go build -o /go-ascii-art

# Expose the application port
ENV port=8080

EXPOSE 8080

# Set the command to run the application
CMD ["/go-ascii-art"]
