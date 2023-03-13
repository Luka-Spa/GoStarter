FROM golang:1.20-bullseye as builder

RUN go install golang.org/dl/go1.20@latest \
    && go1.20 download

WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOAMD64=v3 go build -o ./app -tags timetzdata -trimpath .

FROM scratch
COPY --from=builder /build/app /app

ENTRYPOINT ["/app"]
