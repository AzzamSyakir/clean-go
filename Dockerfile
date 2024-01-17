# Start from golang base image
FROM golang:alpine 

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set working directory di dalam container
WORKDIR /var/www/html