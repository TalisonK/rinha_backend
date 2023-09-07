#First stage: build the executable.
FROM golang:1.20 AS builder

WORKDIR /app

COPY . /app/

# Download dependencies
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server

#second stage: build the image

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /server /server

EXPOSE 3033

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]
