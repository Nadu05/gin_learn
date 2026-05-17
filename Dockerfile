FROM ubuntu:latest
LABEL authors="naduk"

ENTRYPOINT ["top", "-b"]