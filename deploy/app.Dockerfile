FROM golang:1.20-alpine AS build

WORKDIR /src

# main
COPY cmd/app/ ./cmd/app/
COPY cmd/cmd.go ./cmd/cmd.go

# app
COPY app/ ./app/

# service interface
COPY urlservice/proto/ ./urlservice/proto/
COPY urlservice/service.go ./urlservice/service.go

COPY go.mod ./go.mod
RUN go mod tidy

RUN go build -o ./app_bin ./cmd/app

# run

FROM gcr.io/distroless/static-debian11

COPY --from=build /src/app_bin ./

ENTRYPOINT [ "./app_bin" ]
