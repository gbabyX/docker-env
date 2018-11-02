# docker-env
借助docker-compose实现搭建可配置环境
- 安装docker-compose

```
sudo curl -L "https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

- 启动相关容器
```
docker-compose up --build [container]
```
