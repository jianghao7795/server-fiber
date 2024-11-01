#!/bin/bash

docker stop fiber
docker rm fiber
docker rmi fiber

docker build --progress=plain -t fiber .
docker run --name fiber -d -v /root/man/server-fiber/log/:/app/log/ -v /root/man/server-fiber/uploads:/app/uploads -p 3100:3100 fiber
