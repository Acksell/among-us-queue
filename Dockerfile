# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.15.1

WORKDIR /go/src/github.com/Acksell/among-us-queue



# Copy the local package files to the container's workspace.
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the among-us-queue command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install

# Run the among-us-queue command by default when the container starts.
ENTRYPOINT among-us-queue