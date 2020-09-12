# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.15.1

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/Acksell/among-us-queue

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install /go/src/github.com/Acksell/among-us-queue

# Run the outyet command by default when the container starts.
ENTRYPOINT among-us-queue