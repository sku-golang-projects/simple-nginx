FROM alpine:3.19.1

WORKDIR /home/app

COPY ./main /home/app/

CMD "./home/app/main"