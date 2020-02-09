FROM golang:1.13 AS buildenv

LABEL maintainer="Raphael Santo Domingo <raphael.santodomingo@gpmail.com>"

ARG BIN
ARG VERSION
ARG GIT_COMMIT

ENV BIN=${BIN}
ENV VERSION=${VERSION}
ENV GIT_COMMIT=${GIT_COMMIT}

ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN mkdir -p /app

COPY go.* /app/

WORKDIR /app
RUN go mod download
RUN go mod verify

COPY internal internal
COPY cmd cmd

RUN go build \
    -mod=readonly \
    -ldflags "-X github.com/pa3ng/dt-breakfast/internal/buildinfo.Version=$VERSION -X github.com/pa3ng/dt-breakfast/internal/buildinfo.GitCommit=$GIT_COMMIT -X 'github.com/pa3ng/dt-breakfast/internal/buildinfo.BuildTime=$(date -u '+%Y-%m-%d %H:%M:%S')'" \
    -a -o /go/bin/app ./cmd/"$BIN"

FROM scratch
COPY --from=buildenv /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]