# Araj 
Skeleton for microservices

# ENV VARS sample
PORT=8443
APPNAME=araj
DEVELOPMENT=true
CERTFILE=<path>/cert.pem
KEYFILE =<path>/key.pem
ADDRESS=127.0.0.1

#Compile for alpine linux
The Docker file use alpine linux (5.03 MB)
use one of this two option to build the gocode

CGO_ENABLED=1 go build -tags netgo -a -v

or 

CGO_ENABLED=0
 
at the end the docker image will be 14.09 MB
 
# Docker build
cd <path>/araj/src
docker build -t desmax/araj .
 
# Docker run
docker run -it desmax/araj 
-e PORT='8080' \
-e APPNAME='araj' \
-e DEVELOPMENT='true' \
-p 8080:8080 \

 __,   ,_  __,   .  __/
(_/(__/ (_(_/(__/__(_/(_
              _/_
             (/


#Container ip
docker inspect -f '{{.Name}} - {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -aq)
 

