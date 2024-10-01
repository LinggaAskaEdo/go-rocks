FROM golang:1.23.1-alpine3.20 AS golang

RUN apk add -U tzdata
RUN apk --update add ca-certificates --no-cache git

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify
RUN go mod tidy
RUN go generate ./src/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./src/cmd

FROM scratch

COPY --from=golang /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golang /etc/passwd /etc/passwd
COPY --from=golang /etc/group /etc/group

COPY --from=golang /app/.env .
COPY --from=golang /app/etc/cert/* ./etc/cert/
COPY --from=golang /app/.gen/* ./.gen/
COPY --from=golang /app/docs/* ./docs/

COPY --from=golang /server .

EXPOSE 8181

CMD ["/server"]