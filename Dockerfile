FROM golang:1.23.3-alpine3.19 as Develop
WORKDIR /app

RUN apk update && \
apk add --no-cache tzdata

ENV TZ Asia/Bangkok

RUN go install github.com/air-verse/air@latest

COPY ./code/go.mod ./code/go.sum ./
RUN go mod download && go mod verify

CMD ["air", "-c", ".air.toml"]
EXPOSE 8080


FROM golang:1.22.5-alpine3.19 as Nolive
WORKDIR /app

RUN apk update && \
apk add --no-cache tzdata

ENV TZ Asia/Bangkok


COPY ./code/go.mod ./code/go.sum ./
COPY ./code/ .
RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]
EXPOSE 8080


FROM golang:1.22.5-alpine3.19 as Build
WORKDIR /usr/src/app

RUN apk update && \
apk add --no-cache tzdata

ENV TZ Asia/Bangkok

COPY ./code/go.mod ./code/go.sum ./
RUN go mod download && go mod verify

COPY ./code/. .

RUN go build -o main .

FROM golang:1.22.5-alpine3.19 AS Prod
WORKDIR /usr/src/app

RUN apk update && \
apk add --no-cache tzdata

ENV TZ Asia/Bangkok

COPY --from=Build /usr/src/app/main .
COPY ./code/conf/. ./conf/.
CMD ["./main"]

EXPOSE 8080