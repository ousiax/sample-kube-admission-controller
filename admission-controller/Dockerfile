FROM golang:1.17-bullseye
ENV GOPROXY="https://goproxy.cn,direct"
WORKDIR /build

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o admission

FROM debian:bullseye-slim
COPY --from=0 /build/admission /admission
EXPOSE 8080
ENTRYPOINT ["/admission"]
