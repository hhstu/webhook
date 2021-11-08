# shellcheck disable=SC1113
#/bin/bash

# https://www.jianshu.com/p/753478c90049 自行生成证书

cat server.crt | base64 | tr -d '\n'
