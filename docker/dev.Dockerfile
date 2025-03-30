FROM golang:1.23.7-alpine3.21

WORKDIR /app

# Install Air for live reloading
RUN go install github.com/cosmtrek/air@v1.42.0 && \
    ln -s /go/bin/air /usr/local/bin/air  # Symlink air to /usr/local/bin

# Copy dependencies first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

EXPOSE 8000

# Use /bin/sh explicitly to prevent exec format error
CMD ["/bin/sh", "-c", "air"]