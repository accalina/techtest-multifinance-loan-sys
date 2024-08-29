FROM golang:alpine3.18 AS buildStage
LABEL author="github.com/accalina"

ENV GO111MODULE=on CGO_ENABLED=1
WORKDIR /fintech
COPY ./go.mod go.mod
COPY ./go.sum go.sum
COPY ./cmd cmd
COPY ./config config
COPY ./delivery delivery
COPY ./entity entity
COPY ./infra infra
COPY ./repository repository
COPY ./usecase usecase

RUN apk add upx util-linux-dev build-base
RUN go mod download
RUN go test mf-loan/delivery/http/tests mf-loan/repository/tests mf-loan/usecase/tests -v
ENV GO111MODULE=on CGO_ENABLED=0
RUN go build -a -gcflags=all="-l -B" -ldflags "-s -w" -o ./fintech-app cmd/main.go
RUN upx -9 --lzma /fintech/fintech-app
RUN echo 'DB_DSN="root:fintech-password@tcp(mysql:3306)/loan_engine_db?charset=utf8mb4&parseTime=True&loc=Local"' > /fintech/temp.conn

FROM gcr.io/distroless/static:latest AS runtimeStage
WORKDIR /app
COPY --from=buildStage /fintech/fintech-app .
COPY --from=buildStage /fintech/temp.conn .env
ENTRYPOINT ["/app/fintech-app"]