```
步骤 | 作用 | 使用的 GitHub Actions 官方库
Checkout 代码 | 下载代码 | actions/checkout@v4
设置 Go 环境 | 安装 Go | actions/setup-go@v5
登录 DockerHub | 登录凭证 | docker/login-action@v3
Build+Push 镜像 | 构建并推送 Docker 镜像 | docker/build-push-action@v5
```