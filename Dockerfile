FROM golang:1.23.1-alpine

# Install necessary dependencies
RUN apk add --no-cache git bash
# Install Delve dependencies
RUN apk add --no-cache --virtual .build-deps go gcc musl-dev
# Install delve for debugging
RUN go install github.com/go-delve/delve/cmd/dlv@latest
# Remove build dependencies to reduce the final image size after delve installation
RUN apk del .build-deps
# Set environment variables to ensure go binaries are in PATH
ENV PATH="/go/bin:${PATH}"
# Install goose for database migration
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# Install sqlc for generating SQL code
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# Install air for hot reload
RUN go install github.com/air-verse/air@latest
# Copy all code files to the container
COPY . /app
# Set the working directory
WORKDIR /app
# Download all dependencies
RUN go mod download

EXPOSE 8080 2345

CMD [ "air", "-c", ".air.toml" ]