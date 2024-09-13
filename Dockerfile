FROM golang:1.23.1-alpine

# Set environment variables to ensure go binaries are in PATH
ENV PATH="/go/bin:${PATH}"
# Install necessary dependencies
RUN apk add --no-cache git
# Install goose for database migration
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# Install air for hot reload
RUN go install github.com/air-verse/air@latest
# Copy all code files to the container
COPY . /app
# Set the working directory
WORKDIR /app
# Download all dependencies
RUN go mod download

EXPOSE 8080

CMD [ "air", "-c", ".air.toml" ]