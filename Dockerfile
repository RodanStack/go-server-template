FROM golang:1.23.1-alpine

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