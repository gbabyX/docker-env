# docker-env
借助docker-compose实现搭建可配置环境
- 安装docker-compose

```
sudo curl -L "https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```
- 创建.env 文件
```
cp .env_example .env
```
- 配置工作目录
在.env文件中修改PROJECT_SRC为本地的工作目录位置
- 启动相关容器
```
docker-compose up --build [container]
```
