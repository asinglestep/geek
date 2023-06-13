# 运行
```
go run httpserver/main.go --config ./httpserver/etc/config.toml
```

# 生成镜像
```
make httpserver-image
```

# k8s部署
```
kubectl apply -f https://raw.githubusercontent.com/asinglestep/geek/master/httpserver/deployment/deployment.yaml
```