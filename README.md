# freeTranslate

## docker

### /etc/docker/deamon

```json
{
  "builder": {
    "gc": {
      "defaultKeepStorage": "20GB",
      "enabled": true
    }
  },
  "dns": [
    "8.8.8.8"
  ],
  "experimental": false
}
```
`docker build -f dockerfile -t trans:v4 .`
`docker-compose up`

配置文件使用环境变量方法传入

## 改用sqllite作为数据库

文件名`trans.db`