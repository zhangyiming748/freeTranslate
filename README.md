# freeTranslate

读取程序目录下`before.srt`

根据配置文件中的`form`字段设置翻译对应语言字幕到`after.srt`中

# `conf.ini`文件
在百度翻译开放平台[开发者中心](https://fanyi-api.baidu.com/manage/developer)可以找到appid和key,注册使用都是免费的
```ini
[root]
dir = /Users/zen/Downloads/Puppetmaster/Sensual Adventures/srt
;apk add translate-shell
;apt-get install translate-shell
;dnf install translate-shell
;brew install translate-shell
[shell]
;from = ja
from = en
to = zh
proxy = 127.0.0.1:7890
```

# 改用sqllite作为数据库

文件名`trans.db`