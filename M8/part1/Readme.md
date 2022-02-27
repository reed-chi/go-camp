- 优雅启动
```
	startupProbe:
	  httpGet:
		path: /ping
		port: 8080 
		scheme: HTTP
		initialDelaySeconds: 5
		periodSeconds: 10
		successThreshold: 1
		timeoutSeconds: 3  
```
- 优雅终止
```
	terminationGracePeriodSeconds: 30 
```

- 资源需求和 QoS 保证
```
resources:
	limits:
	  cpu: '500m'
	  memory: 1Gi
	requests:
	  cpu: 500m
	  memory: 1Gi   
```
- 探活
```
livenessProbe:
  httpGet:
    path: /ping
    port: 8080 
    scheme: HTTP
    initialDelaySeconds: 5
    periodSeconds: 10
    successThreshold: 1
    timeoutSeconds: 3
```
- 日常运维需求，日志等级
```
env:
  - name: aliyun_logs_webapp
    value: "stdout"	
```

```
volumeMounts:
- mountPath: /data
  name: volume-alicloud-nas-pvc
volumes:
- name: volume-alicloud-nas-pvc
persistentVolumeClaim:
  claimName: alicloud-nas-pvc
```
- 配置和代码分离
```
配置存configmap或是配置中心
```
