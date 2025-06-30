FROM golang:1.24.2-alpine3.21 AS build
WORKDIR /app
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN go build -o /app/paracad .

# FROM alpine:3.21
# RUN apk add --no-cache openscad
FROM openscad/openscad:egl
COPY --from=build /app/paracad /app/paracad
COPY images /app/images
COPY not-found.png /app/images
EXPOSE 8081
CMD ["/app/paracad"]
