#!/bin/bash
export MINIO_ROOT_USER=minio
export MINIO_ROOT_PASSWORD=miniostorage
minio server $(pwd)/data &
