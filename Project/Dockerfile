FROM golang:alpine AS builder

RUN apk update && \
    apk add git && \
    apk add build-base upx

WORKDIR /src/app
COPY . .

RUN go build -o /go/bin/app cmd/main.go
RUN upx /go/bin/app

# spep 2

FROM alpine


RUN apk update && apk add --no-cache vips-dev
COPY --from=builder /go/bin/app /go/bin/app
# COPY --from=builder /src/app/cmd/config.yaml /go/bin

# ENTRYPOINT ["go/bin/app", "-configpath=/go/bin/config.yaml"]
ENTRYPOINT ["go/bin/app"]
###############################
# FROM golang

# WORKDIR /go/bin/app
# COPY . /go/bin/app
 