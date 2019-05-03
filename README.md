# Go playground [![Build Status](https://travis-ci.org/ppworks/go-playground.svg?branch=master)](https://travis-ci.org/ppworks/go-playground)

## Docker

docker build . -t go-playground
docker run -it --rm -p 3000:3000 -v $PWD:/go/src/app go-playground

or

docker-compose up

## Webpack Dev Server

yarn run start
