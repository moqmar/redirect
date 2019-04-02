FROM golang AS build

ADD . /go/src/github.com/moqmar/redirect
RUN go get github.com/moqmar/redirect
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static" -s -w' -o /tmp/redirect github.com/moqmar/redirect

FROM scratch
COPY --from=build /tmp/redirect /redirect
EXPOSE 80
CMD ["/redirect"]
