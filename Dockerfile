FROM golang:1.18-alpine
RUN go install github.com/cosmtrek/air@latest
WORKDIR /app
CMD . ./
RUN go mod download
RUN go mod tidy
CMD ["air"]
