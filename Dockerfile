FROM golang:alpine as builder
# ARG SOURCE_LOCATION=/app
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 3000
CMD ["./main"]
