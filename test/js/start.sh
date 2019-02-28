#!/usr/bin/env bash

#grpcc -i -p ../../protobuf/imgtrip.proto -a localhost:50050 --exec post.js
#grpcc -i -p ../../protobuf/imgtrip.proto -a localhost:50050 --exec comment.js
#grpcc -i -p ../../protobuf/imgtrip.proto -a localhost:50050 --exec image.js
grpcc -i -p ../../protobuf/imgtrip.proto -a localhost:50050 --exec album.js
#grpcc -i -p ../../protobuf/imgtrip.proto -a localhost:50050 --exec review.js