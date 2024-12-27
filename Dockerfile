FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN go build -o ./blober .

FROM alpine
COPY --from=build /app/blober .
RUN apk add libc6-compat 
EXPOSE 8978
CMD ["./blober"]
