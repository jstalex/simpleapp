FROM golang:alpine AS build

WORKDIR /simpleapp

ENV GO111MODULE=on

COPY . .
RUN go build -o main cmd/app/main.go

EXPOSE 8080

FROM scratch

COPY --from=build simpleapp/main /bin/main

ENTRYPOINT ["/bin/main"]