1. 接收客户端 request，并将 request 中带的 header 写入 response header

    `curl -I -H 'Geek:geek' 127.0.0.1:8080/header`

2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header

 `curl -I -H 'Geek:geek' 127.0.0.1:8080/version`


3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出

  ```
  127.0.0.1 - [Tue, 04 Jan 2022 22:10:00 CST] "HEAD /version HTTP/1.1 200 3.402µs "curl/7.61.1" "
  127.0.0.1 - [Tue, 04 Jan 2022 22:10:14 CST] "HEAD /version HTTP/1.1 200 3.364µs "curl/7.61.1" "
  127.0.0.1 - [Tue, 04 Jan 2022 22:10:31 CST] "GET /healthz HTTP/1.1 200 46.737µs "curl/7.61.1" "
  ```
  
4. 当访问 localhost/healthz 时，应返回 200

  `curl   127.0.0.1:8080/healthz`
