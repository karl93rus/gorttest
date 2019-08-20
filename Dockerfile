FROM golang:1.12-buster

RUN apt-get update && apt-get install cron -y

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get github.com/lib/pq

COPY upd-cron /etc/cron.d/upd-cron
RUN chmod a+x /etc/cron.d/upd-cron
RUN crontab /etc/cron.d/upd-cron

RUN go build -o server .

# CMD ["cron && ./server"]
CMD crontab -l && cron && ./server
