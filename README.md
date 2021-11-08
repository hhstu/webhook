# k8s webhook demo

./secret/crt.sh 是为了 https 自签证书的脚本，替换 lc.com 为自己的domain；

./k8s/webhook.yaml 是提交给k8s的配置，需要改一下 url 的地址


参考：

https://kubernetes.io/zh/docs/reference/access-authn-authz/extensible-admission-controllers/
