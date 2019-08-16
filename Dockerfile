FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN rm Dockerfile

RUN cp ./Docker/Dockerfile .

CMD tar cvzf - .
