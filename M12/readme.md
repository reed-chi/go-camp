# 模块12作业

把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：

1. 如何实现安全保证
```
ubuntu@VM-100-17-ubuntu:~$ kubectl  -n securesvc  get pods
NAME                          READY   STATUS    RESTARTS   AGE
httpserver-5b5fc75b98-bbjnj   2/2     Running   2          2d
ubuntu@VM-100-17-ubuntu:~$ kubectl  -n istio-system  get svc
NAME                   TYPE           CLUSTER-IP       EXTERNAL-IP      PORT(S)                                                                      AGE
istio-egressgateway    ClusterIP      172.20.253.51    <none>           80/TCP,443/TCP                                                               3d2h
istio-ingressgateway   LoadBalancer   172.20.253.132   119.28.141.248   15021:30039/TCP,80:32706/TCP,443:31433/TCP,31400:30848/TCP,15443:31702/TCP   3d2h
istiod                 ClusterIP      172.20.253.22    <none>           15010/TCP,15012/TCP,443/TCP,15014/TCP                                        3d2h
ubuntu@VM-100-17-ubuntu:~$
ubuntu@VM-100-17-ubuntu:~$
ubuntu@VM-100-17-ubuntu:~$ curl --resolve httpsserver.cncamp.io:443:$INGRESS_IP https://httpsserver.cncamp.io/healthz -v -k
* Added httpsserver.cncamp.io:443:172.20.253.132 to DNS cache
* Hostname httpsserver.cncamp.io was found in DNS cache
*   Trying 172.20.253.132...
* TCP_NODELAY set
* Connected to httpsserver.cncamp.io (172.20.253.132) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
```
2. 七层路由规则
```
root@VM-100-17-ubuntu:~# kubectl  -n simple  get pods
NAME                                READY   STATUS    RESTARTS   AGE
nginx-deployment-6799fc88d8-grvrz   1/1     Running   0          4m21s
simple-7697f7dbdd-vgl7m             1/1     Running   0          31s
root@VM-100-17-ubuntu:~# kubectl  -n simple  get svc
NAME     TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
nginx    ClusterIP   172.20.254.94   <none>        80/TCP    4m25s
simple   ClusterIP   172.20.255.17   <none>        80/TCP    35s
root@VM-100-17-ubuntu:~# kubectl  -n simple  get vs
NAME     GATEWAYS     HOSTS                  AGE
simple   ["simple"]   ["simple.cncamp.io"]   4m29s
root@VM-100-17-ubuntu:~# kubectl  -n simple  get gw
NAME     AGE
simple   4m33s
root@VM-100-17-ubuntu:~# curl -H "Host: simple.cncamp.io" $INGRESS_IP/simple/hello
hello [stranger]
===================Details of the http request header:============
X-Envoy-Internal=[true]
X-Request-Id=[a7d44f84-7993-9bed-adc3-880a6b8594d5]
X-B3-Traceid=[68290d80c7e4884b384a12522c52a4d5]
User-Agent=[curl/7.58.0]
X-Forwarded-For=[172.19.100.17]
X-Envoy-Decorator-Operation=[simple.simple.svc.cluster.local:80/simple/hello]
X-Envoy-Original-Path=[/simple/hello]
Accept=[*/*]
X-Envoy-Peer-Metadata=[ChF5]
X-Envoy-Peer-Metadata-Id=[router~172.20.0.18~istio-ingressgateway-c45ccd4dc-n2wjb.istio-system~istio-system.svc.cluster.local]
X-B3-Spanid=[384a12522c52a4d5]
X-B3-Sampled=[1]
X-Forwarded-Proto=[http]
X-Envoy-Attempt-Count=[1]
root@VM-100-17-ubuntu:~# curl -H "Host: simple.cncamp.io" $INGRESS_IP/nginx
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

3. 考虑 open tracing 的接入
```
ubuntu@VM-100-17-ubuntu:~$ kubectl  -n istio-system  get vs
NAME         GATEWAYS              HOSTS                               AGE
tracing-vs   ["tracing-gateway"]   ["tracing.124.156.122.11.nip.io"]   14h
```
[tracing](https://github.com/reed-chi/go-camp/blob/main/M12/tracing.png)
