FROM golang:1.24.2-alpine3.21 AS build
WORKDIR /app
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN go build -o /app/paracad .

FROM alpine:3.21
RUN apk add --no-cache openscad
COPY --from=build /app/paracad /app/paracad
EXPOSE 8081
CMD ["/app/paracad"]
