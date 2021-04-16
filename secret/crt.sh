# shellcheck disable=SC1113
#/bin/bash
openssl req -x509 -sha256 -newkey rsa:4096 -keyout ca.key -out ca.crt -days 3560 -nodes -subj '/CN=My Cert Authority'

echo "生成用上述 ca 签署的 server 证书"
### 多域名、泛域名 配置 -subj '/CN=lc.com/CN=*.xxx.com/CN=172.20.76.21'
openssl req -new -newkey rsa:4096 -keyout server.key -out server.csr -nodes -subj '/CN=lc.com'
openssl x509 -req -sha256 -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt


cat server.crt | base64 | tr -d '\n'
