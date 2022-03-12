1. 为 HTTPServer 添加 0-2 秒的随机延时；
2. 为 HTTPServer 项目添加延时 Metric；
3. 将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
4. 从 Promethus 界面中查询延时指标数据；
![图片alt](https://github.com/reed-chi/go-camp/blob/main/M10/prometheus-snapshot.png "Prometheus截图")
5. 创建一个 Grafana Dashboard 展现延时分配情况。
6.  ![图片alt](https://github.com/reed-chi/go-camp/blob/main/M10/grafana-snapshot.png "Prometheus截图")


