# Set the base image
FROM golang:latest

# Set metadata labels
LABEL maintainer = "Rami, Hasan, Hussain"
LABEL descriptino = "ASCII art web Docker image"

# Set the working directory
WORKDIR /app

# Copy the application code
COPY . .

# Build the application
RUN go build -o PICASO cmd/asciiartweb/main.go

# Expose the application port
ENV port=8080

EXPOSE 8080

# Set the command to run the application
CMD ["/app/PICASO"]