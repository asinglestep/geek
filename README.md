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

# istio ingress gateway部署
```
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.com' -keyout cncamp.com.key -out cncamp.com.crt
kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.com.key --cert=cncamp.com.crt

kubectl create ns tracing
kubectl label ns tracing istio-injection=enabled
kubectl apply -f istio.yaml
```

## 测试
```
kubectl get svc -n istio-system istio-ingressgateway
export INGRESS=10.99.43.20
curl --resolve httpserver.cncamp.com:443:$INGRESS https://httpserver.cncamp.com/random -v -k
```