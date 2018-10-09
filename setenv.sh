#!/bin/bash
# MEMO: Right one will override the left one
LOADED_ENV_ORDER='.env'
echo "LOADED_ENV_ORDER:"${LOADED_ENV_ORDER}
export $(grep -v -h '^#' $(echo ${LOADED_ENV_ORDER} | xargs) | xargs)
echo "Successfully set environment variables."
