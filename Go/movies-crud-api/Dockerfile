FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get .
RUN go build -o main .

EXPOSE 8000
CMD [ "/app/main" ]