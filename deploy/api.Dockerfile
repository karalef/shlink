FROM golang:1.20-alpine AS build

WORKDIR /src

# main
COPY cmd/service/ ./cmd/service/
COPY cmd/cmd.go ./cmd/

# storage
COPY model/ ./model
COPY repo/ ./repo

# service
COPY urlservice/ ./urlservice

COPY go.mod ./go.mod
RUN go mod tidy

RUN go build -o ./api ./cmd/service

# run

FROM gcr.io/distroless/static-debian11

COPY --from=build /src/api ./

ENTRYPOINT [ "./api" ]
