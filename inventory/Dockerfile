FROM golang as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o inventory/inventory inventory/main.go


FROM scratch

COPY --from=builder /app/inventory /app/

ENTRYPOINT ["/app/inventory"]