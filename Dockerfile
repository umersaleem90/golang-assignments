FROM golang:1.16-alpine

# Maintainer information
LABEL maintainer="umer.saleem@tintash.com"

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .

RUN go build
EXPOSE 3000
CMD ["./todo-backend"]