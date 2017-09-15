FROM frolvlad/alpine-glibc

RUN mkdir /app
ADD go_mytodo /app/

CMD ["/app/go_mytodo"]
