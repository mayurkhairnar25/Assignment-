FROM golang:1.14-alpine AS build
LABEL maintainer <"Quique<mayurkhairnar325@gmail.com">
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV PORT 8088
RUN go build
CMD ["./Docker2"]

