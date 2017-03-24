# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/mattherman/url-shortener

# Get all necessary dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/garyburd/redigo/redis

# Build the url-shortener command inside the container.
RUN go install github.com/mattherman/url-shortener

# Run the url-shortener command by default when the container starts.
ENTRYPOINT /go/bin/url-shortener

# Document that the service listens on port 8080.
EXPOSE 8080