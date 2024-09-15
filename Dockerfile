# Build stage
FROM golang:1.22.3-alpine3.19 AS build
WORKDIR /app
COPY . .
COPY go.mod ./
RUN go mod download && go mod verify
RUN go build -o main .

# Final stage
FROM scratch
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]