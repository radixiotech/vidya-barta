FROM golang:1.22.4-bullseye AS build

WORKDIR /app

COPY apps/vb-backend/go.mod .
COPY apps/vb-backend/go.sum .

RUN go mod download && go mod verify

COPY apps/vb-backend/apps /app/apps
COPY apps/vb-backend/business /app/business
COPY apps/vb-backend/foundation /app/foundation

RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o app -a -ldflags="-s -w" -installsuffix cgo apps/vb-api/main.go

RUN apt-get update && apt-get install -y --no-install-recommends \
  upx \
  build-essential \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*
RUN upx --ultra-brute -qq app && upx -t app

FROM scratch

WORKDIR /app

COPY --from=build /app/app /app

ENTRYPOINT [ "/app/app" ]