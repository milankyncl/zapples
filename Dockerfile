FROM node:19.3.0-alpine3.17 AS ui-builder

WORKDIR /src

COPY ui/ .

RUN yarn --frozen-lockfile

RUN yarn build


FROM golang:1.19.4-alpine3.17 AS api-builder

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk add --no-cache \
    gcc \
    musl-dev

WORKDIR /src

COPY . .

RUN go build -o /build -ldflags='-w -extldflags "-static"' main.go


FROM alpine:3.17

ENV ZAPPLES_BASE_PATH="/usr/share/zapples"

RUN mkdir -p $ZAPPLES_BASE_PATH && chown 499:499 -R $ZAPPLES_BASE_PATH

USER 499:499

COPY migrations/ "${ZAPPLES_BASE_PATH}/migrations"

COPY --from=ui-builder /src/dist/ "${ZAPPLES_BASE_PATH}/ui/dist"
COPY --from=api-builder /build /usr/local/bin/zapples

ENTRYPOINT ["/usr/local/bin/zapples"]
