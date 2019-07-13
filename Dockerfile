# Start from golang:1.12-alpine base image
FROM golang:1.12-alpine

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/ugniusin/watchme

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Set GO111MODULE off, because running the module within GOPATH
ENV GO111MODULE=off

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Install package to watch changes
RUN go get github.com/oxequa/realize

# This container exposes port 8090 to the outside world
EXPOSE 8090

# Run the executable
CMD ["realize", "start"]
