# docker build -t nginx-http3 .

FROM archlinux:base-devel

RUN pacman -Syu --noconfirm && \
        pacman -S --noconfirm nginx-mainline

RUN mkdir /qlog

COPY nginx.conf /etc/nginx/nginx.conf
COPY localhost.key /etc/nginx/certs/localhost.key
COPY localhost.crt /etc/nginx/certs/localhost.crt
COPY index.html /etc/nginx/html/index.html

EXPOSE 443

CMD ["/usr/bin/nginx"]