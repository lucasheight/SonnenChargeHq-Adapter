#!/bin/bash
 mkdir ./dist/docker --parent
 docker save -o .//dist//docker//go-sonnen-chargehq-service.tar go-sonnen-chargehq-service:prod