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
./run.sh demo

# 监听8082
./run.sh demo -addr=127.0.0.1:8082
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
- [ ] 签名机制实现
- [ ] 细节优化

## 原理
### client one
1. 客户端页面请求发现会话过期，需要登录
2. 跳转至sso中心登陆页面，填写用户信息登陆
3. sso系统登陆后，在sso系统域名下保存会话
4. sso保存会话完成，携带token跳转至客户端页面
5. 客户端根据token请求sso，获取用户信息
6. 客户端获取用户信息后，保存至本地会话信息
7. 客户端每次请求都访问sso系统进行检查

### client two
1. 其他实现sso客户端的模块，访问页面的时候发现会话过期
2. 跳转至sso系统登陆页面，但是sso系统已登陆
3. sso系统携带token跳转至客户端页面
4. 重复上面5-7