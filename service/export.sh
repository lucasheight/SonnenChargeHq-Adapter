#!/bin/bash
 mkdir ./dist/docker --parent
 docker save -o .//dist//docker//sonnen-chargehq-service.tar sonnen-chargehq-service:1.0.0