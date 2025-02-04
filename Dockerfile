FROM golang:1.23.5-alpine

ARG DB_USER=myuser
ARG DB_PASSWORD=mypassword

ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_NAME=mydatabase

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o service-account .

EXPOSE 8080

CMD ["./service-account"]
