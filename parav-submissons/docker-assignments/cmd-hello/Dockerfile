FROM  golang:1.16-alpine

WORKDIR /src

COPY main.go /src/
ENTRYPOINT [ "/src/main", "-input=Hello World!!!"]
RUN go build main.go

CMD ["./main"]