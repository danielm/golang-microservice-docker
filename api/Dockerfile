FROM golang:1.14.1

WORKDIR /usr/src/app

COPY ./app ./

# ADD ./app ./

RUN go get github.com/gorilla/mux

RUN go build -o "/usr/local/bin/api-main" ./src

# EXPOSE 8080

CMD ["api-main"]
