FROM ubuntu
LABEL authors="aaditya"

RUN mkdir /app
WORKDIR /app
COPY build/run .

ENTRYPOINT ./run