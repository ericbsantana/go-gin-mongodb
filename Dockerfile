FROM golang:latest

LABEL maintainer="Eric Santana <bitencourteric@hotmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin && chmod +x /go/bin/air

EXPOSE 8080

CMD ["/go/bin/air"]