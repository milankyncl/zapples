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

ENV ZAPPLES_PATH="/usr/share/zapples"

RUN mkdir -p $ZAPPLES_PATH && chown 499:499 -R $ZAPPLES_PATH

USER 499:499

ENV ZAPPLES_DB_PATH="$ZAPPLES_PATH/local.sqlite"
ENV ZAPPLES_MIGRATIONS_PATH="$ZAPPLES_PATH/migrations"
ENV ZAPPLES_UI_PATH="$ZAPPLES_PATH/ui"

COPY migrations/ $ZAPPLES_MIGRATIONS_PATH

COPY --from=ui-builder /src/dist/ $ZAPPLES_UI_PATH
COPY --from=api-builder /build /usr/local/bin/zapples

ENTRYPOINT ["/usr/local/bin/zapples"]
