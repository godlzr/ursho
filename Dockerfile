FROM golang:latest

ADD . /go/src/github.com/Ziyang2go/ursho/

WORKDIR /go/src/github.com/Ziyang2go/ursho/

RUN go get && go build

RUN rm Dockerfile

RUN cp ./Docker/Dockerfile .

CMD tar cvzf - .
