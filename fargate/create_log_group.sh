#!/bin/bash
if [ -z "${AWS_PROFILE}" ]
then
  echo "$0: you need to specify AWS_PROFILE=..."
  exit 1
fi
aws logs create-log-group --log-group-name sds-logs
