FROM d.quantibio.com/alpine:3.3

ADD k
EXPOSE 443
EXPOSE 8080

RUN ls /bin