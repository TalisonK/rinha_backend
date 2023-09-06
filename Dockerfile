# syntax=docker/dockerfile:1

FROM golang:1.20 aS builder

WORKDIR /app

# Copy Go module files
COPY go.* ./

# Download dependencies
RUN go mod download

# Copy source files
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /bin/server .

FROM alpine:latest AS final

EXPOSE 8080

# Copy files from builder stage
COPY --from=build /bin/server /bin/
COPY ./config.toml /data/config.toml



ENTRYPOINT [ "/bin/server" ]
