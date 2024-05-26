FROM --platform=$BUILDPLATFORM golang:alpine as builder
WORKDIR /go/src/app

RUN apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS TARGETARCH
# Static build required so that we can safely copy the binary over.
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o ./bin/server -ldflags '-extldflags "-static"' ./cmd/server/main.go

FROM scratch
COPY --from=builder /go/src/app/bin/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
