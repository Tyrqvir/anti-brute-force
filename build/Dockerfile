FROM golang:1.19-alpine as build

ARG BIN_FILE
ARG APP_FOLDER

ENV BIN_FILE $BIN_FILE
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Cache layer with modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# Building a static Go binary (no CI dependencies),
# otherwise it won't work in alpine image.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} ./cmd/${APP_FOLDER}

# Make thin image
FROM alpine:latest

ARG CONFIG_FILE_EXTERNAL

ARG BIN_FILE
ENV BIN_FILE $BIN_FILE
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/app/config.toml
COPY ${CONFIG_FILE_EXTERNAL} ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}
