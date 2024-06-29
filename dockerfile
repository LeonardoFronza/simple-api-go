FROM golang:latest AS BUILD

WORKDIR /app

COPY ./App/go.mod ./App/go.sum ./
RUN go mod download

COPY ./App/*.go  ./
COPY ./App/docs ./docs

RUN CGO_ENABLED=0 GOOS=linux go build -o ApiRapazes

FROM alpine:latest AS RELEASE
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=BUILD /app .

EXPOSE 8080
RUN chmod a+x ./ApiRapazes

CMD ["./ApiRapazes"]