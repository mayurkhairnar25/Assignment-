FROM golang:latest as builder
LABEL maintainer="Quique<mayurkhairnar325@gmail.com>"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
ENV PORT 2020
CMD ["./Multistage"]