# sso实践

## 使用
1. 配置hosts
```etc
127.0.0.1 client.one.com
127.0.0.1 client.two.com
127.0.0.1 sso.com
```

2. 初始化mysql数据库和表
```shell
./run.sh init
```

3. 编译主程序，默认监听8099端口
4. 启动两个客户端程序
```shell
# 默认监听8081
./run.sh client

# 监听8082
./run.sh client -addr=127.0.0.1:8082
```

4. 浏览器访问
```
client.one.com:8081/index
client.two.com:8082/index
```

**默认登录信息：**
```
email: admin@gmail.com
password: admin
```

## TODO
- [X] 登录逻辑
- [X] 退出逻辑
- [ ] 签名机制
- [ ] 引入redis做svc接口
- [X] sso中心后台
