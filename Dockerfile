FROM golang:1.15.5
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]