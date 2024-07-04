FROM golang:1.22.4 as builder

ADD . /build/

WORKDIR /build

RUN go mod download && go mod verify

RUN go build -v -o ../cudos-task cudos-task/cmd/cudos-task

FROM gcr.io/distroless/base

COPY --from=builder /cudos-task /cudos-task

ENTRYPOINT ["./cudos-task"]
