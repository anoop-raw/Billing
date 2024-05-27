# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Install wait-for-it
RUN apt-get update && apt-get install -y wait-for-it

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-billing ./cmd

# Expose port 6000 to the outside world
EXPOSE 6000

# Command to run the executable with wait-for-it
CMD ["wait-for-it", "db:5432", "--", "/docker-billing"]
