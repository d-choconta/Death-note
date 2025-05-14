FROM golang:1.24


RUN apt-get update && apt-get install -y netcat-openbsd


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .


RUN go build -o main .


EXPOSE 8080
COPY wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh
CMD ["/wait-for.sh", "db", "./main"]

