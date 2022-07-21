FROM golang:1.18.4-alpine3.16

WORKDIR /ossinspector
COPY . .
RUN go install

RUN go build main/inspector.go
ENTRYPOINT [ "./inspector" ]
