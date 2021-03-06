FROM debian:buster-slim

ADD ./bin/snowflake-grpc-server /
ADD ./config.json.example /config.json

RUN chmod 755 /snowflake-grpc-server

ENTRYPOINT [ "/snowflake-grpc-server" ]