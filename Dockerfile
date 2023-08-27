# Build the application from source
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY . .

RUN go build -o /shortly

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /shortly /shortly

EXPOSE 8080

CMD ["/shortly"]
