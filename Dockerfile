FROM golang:1-alpine

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD [ "/app/main" ]