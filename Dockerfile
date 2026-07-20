### Etapa 1: Compilación

FROM golang:1.26-alpine AS builder

## Instala git y certificados por si tu app los necesita
RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /app

## Copiamos los archivos del modulo y descargamos dependencias
COPY go.mod ./


RUN go mod download

## Copiamos el resto del codigo fuente
COPY . .

## Compilamos la aplicacion
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o mi-api .

FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /root/

## Copiamos el binario compilado desde la etapa de construcción
COPY --from=builder /app/mi-api .

EXPOSE 8080

CMD ["./mi-api"]