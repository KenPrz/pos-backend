FROM golang:1.24-alpine AS development

WORKDIR /app

ENV CGO_ENABLED=0
ARG AIR_VERSION=v1.61.7

RUN apk add --no-cache ca-certificates git \
    && go install github.com/air-verse/air@${AIR_VERSION}

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]
