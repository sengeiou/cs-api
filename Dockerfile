FROM golang:1.18 as build

WORKDIR /api

COPY . .

#RUN go build -o api main.go

# 去除符号表和调试信息 & 使用 upx 壓縮
RUN go build -ldflags="-s -w" -o server main.go && upx -9 server

FROM gcr.io/distroless/base

WORKDIR /api

COPY --from=build /api/server /api/server
COPY --from=build /api/config.yaml /api/config.yaml

CMD ["./server", "server"]