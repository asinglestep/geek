# 运行
```
go run httpserver/main.go --config ./httpserver/etc/config.toml
```

# 接口文档
## swag安装
```
go install github.com/swaggo/swag/cmd/swag@latest
```

## 生成接口文档
```
make generate 
```

## 查看接口文档
```
http://127.0.0.1:8080/swagger/index.html
```

# 生成镜像
```
make httpserver-image
```

# k8s部署
```
kubectl apply -f https://raw.githubusercontent.com/asinglestep/geek/master/httpserver/deployment/deployment.yaml
```

# prometheus监控
```
httpserver/prometheus/httpserver.json
```