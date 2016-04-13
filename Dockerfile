FROM gliderlabs/alpine
MAINTAINER Massimiliano Dessi (@desmax74)
WORKDIR /app
ADD src/desmax/araj/araj /app/araj
EXPOSE 8080
ENTRYPOINT ["/app/araj"]