FROM golang as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway/gateway gateway/main.go


FROM scratch

COPY --from=builder /app/gateway /app/

ENTRYPOINT ["/app/gateway"]