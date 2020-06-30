FROM golang:1.14-alpine as build-env
WORKDIR /verbugata-api
RUN apk update && apk add --no-cache gcc musl-dev git bash
COPY ./src .
RUN go mod download
RUN ls -al .
RUN go build -ldflags '-w -s' -a -o ./bin/app ./app.go

FROM alpine
RUN apk update && apk add --no-cache bash
COPY --from=build-env /verbugata-api/bin/app /verbugata-api/
EXPOSE 8080
CMD ["/verbugata-api/app"]
