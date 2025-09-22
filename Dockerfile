FROM golang:1.24-alpine3.22 as builder

ARG TOOL_VERSION

WORKDIR /src

RUN apk add git
RUN GO111MODULE=on go install github.com/mgechev/revive@v${TOOL_VERSION}

ADD . .
RUN go build -o bin/codacy-gorevive

FROM alpine:3.22

ENV PATH="/go/bin:${PATH}"
COPY --from=builder /go /go
COPY --from=builder /src/bin /dist/bin
COPY docs/ /docs/

RUN adduser -u 2004 -D docker
RUN chown -R docker:docker /docs

CMD [ "/dist/bin/codacy-gorevive" ]
