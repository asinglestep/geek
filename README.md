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
kubectl apply -f https://github.com/asinglestep/geek/httpserver/deployment/deployment.yaml
```