#!/bin/bash

. ./config.sh

docker run -i -t -h ${GUEST_HOST_NAME} -p ${HOST_FORWARDING_PORT}:${GUEST_FORWARDING_PORT} -v ${HOST_MOUNT_PATH}:${GUEST_MOUNT_PATH} ${REPOSITORY_NAME} /bin/bash