# Build a Docker image for the application using an Alpine Linux image
FROM alpine:latest

# Update the package repository
RUN apk update

# Copy the binary file from your host to your current location.
COPY /bin/GuacAlert_linux .

# Command to run the executable
ENTRYPOINT ["./GuacAlert_linux"]