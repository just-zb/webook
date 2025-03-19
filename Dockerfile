FROM ubuntu:20.04
LABEL authors="michael"
COPY webook /app/webook
WORKDIR /app
CMD ["/app/webook"]