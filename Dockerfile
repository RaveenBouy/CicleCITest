FROM golang:latest

RUN mkdir /app 

ADD . /app/ 

WORKDIR /app/RestAPI

CMD ["./RestAPI"]