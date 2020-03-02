FROM golang

LABEL key="MINT"

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN mkdir /dim-edge-node
WORKDIR /dim-edge-node
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

EXPOSE 9090
EXPOSE 9000

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.BuildEnv=prod" main.go
