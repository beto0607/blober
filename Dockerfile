FROM --platform=$BUILDPLATFORM golang:alpine AS build
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ./blober .

FROM alpine
COPY --from=build /app/blober .
RUN apk add libc6-compat 
EXPOSE 8978
CMD ["./blober"]
