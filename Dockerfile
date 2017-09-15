FROM frolvlad/alpine-glibc

RUN mkdir /app
COPY go_mytodo /app/
COPY ./config /app/config

CMD ["/app/go_mytodo"]
