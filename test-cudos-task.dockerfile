FROM golang:1.22.4

ADD . /build/

WORKDIR /build

RUN go mod download && go mod verify

# Install golangci-lint v1.59.1
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1

# Command to run the linter and the tests
CMD echo "Running linter, please wait..." ; golangci-lint --timeout 5m --verbose run ./... ; go test -a ./... -cover -covermode=atomic -timeout 2m -p 1 -parallel 1 ./...
