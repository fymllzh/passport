# sso实践

## 主要依赖
* [gin](https://github.com/gin-gonic/gin)
* [gorm](https://gorm.io/zh_CN/)
* [go-redis](https://redis.uptrace.dev/)

## 使用
1. 配置hosts
```etc
127.0.0.1 client.one.com
127.0.0.1 client.two.com
127.0.0.1 sso.com
```

2. 调整`conf/app.ini`，初始化mysql数据库和表
* 创建数据库
  
`CREATE DATABASE IF NOT EXISTS passport DEFAULT CHARACTER SET utf8`

* 初始化表结构
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

5. 直接登录sso中心
```
sso.com:8099
```

## TODO
- [X] 登录逻辑
- [X] 退出逻辑
- [X] 签名机制
- [X] 引入redis做svc接口
- [ ] IP白名单
- [X] 登录错误次数限制
- [ ] 后台功能列表
- [ ] 日志处理
- [X] sso中心后台

## 对接流程
### svc接口
* 请求方式
  * POST
  * `Content-Type: application/x-www-form-urlencoded`


* 请求参数

```json
{
  "domain": "client.one.com",
  "timestamp": "1623680856",
  "token": "6bc4931890225677da85a1cf05ce0fc0",
  "sign": "6BC4931890225677DA85A1CF05CE0FC0"
}
```

* sign签名算法
```php
$s = [
    'domain' => 'client.one.com',
    'token' => '6bc4931890225677da85a1cf05ce0fc0',
    'timestamp' => '1623680856',
    'sign' => '6BC4931890225677DA85A1CF05CE0FC0',
];

// 按key排序
ksort($s);

// 拼接内容
$str = implode("", $s);

// 末尾拼接密钥
$str .= "123456";

echo strtoupper(md5($str)), PHP_EOL;
```
