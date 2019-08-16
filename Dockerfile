# Start from golang:1.12-stretch base image
FROM golang:1.12-stretch

# Add apk dependencies
RUN apt-get update && apt-get install -y git openssl

# Add Dockerize to wait for MySQL
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/ugniusin/watchme-backend

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

# Install package to run migrations
RUN go get -v github.com/rubenv/sql-migrate/...

# This container exposes port 8090 to the outside world
EXPOSE 8090

# Run the executable
CMD dockerize -wait tcp://mysql:3306 -timeout 60m realize start
