FROM golang:1.13.5 as builder

RUN go get -u github.com/mgechev/revive

FROM golang:1.13.5-alpine

COPY --from=builder /go /go
COPY bin /dist/bin
COPY docs/ /docs/

RUN adduser -u 2004 -D docker
RUN chown -R docker:docker /docs

CMD [ "/dist/bin/codacy-gorevive" ]
