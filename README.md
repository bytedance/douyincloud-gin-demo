# douyincloud-gin-demo
本项目是抖音云平台基于go语言gin框架的开发模版，模版通过使用Redis和MongoDB实现了简单的hello-world功能。\
抖音云平台支持基于Git代码和Docker镜像部署两种方式。其中，Dockerfile文件可以参考本项目中的Dockerfile文件。
部署在抖音云平台的服务日志需要重定向到标准输出，并在抖音云平台日志功能中查看。

## 目录结构
~~~
.
├── Dockerfile              Dockerfile文件
├── Readme.md               Readme文件
├── go.mod                  go.mod文件
├── go.sum                  go.sum文件
├── main.go                 主函数入口
├── run.sh                  容器运行时启动文件
└── service                 业务逻辑目录
    └── service.go          业务逻辑文件
~~~

## 请求方法
前往抖音云托管平台「调试」功能界面，进行请求调试。

## API说明
### `GET /api/get_open_id`
获取openid，当未绑定用户时会返回获取openid失败的错误


### 响应结果
```json
{
    "err_no": 0,
    "err_msg": "success",
    "data": "719f****-****-4c**-a0**-*********"
}
```

### `POST /api/text/antidirt`
云调用示例，调用抖音开放平台的OpenApi进行脏词检测

### 请求参数
- `content`:`string` 待检测的内容

### 响应结果
```json
{
    "err_no": 0,
    "err_msg": "success",
    "data": ""
}
```

## License

This project is licensed under the [Apache-2.0 License](LICENSE).

