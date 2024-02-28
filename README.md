# freeTranslate

读取程序目录下`before.srt`

根据配置文件中的`form`字段设置翻译对应语言字幕到`after.srt`中

# `conf.ini`文件
在百度翻译开放平台[开发者中心](https://fanyi-api.baidu.com/manage/developer)可以找到appid和key,注册使用都是免费的
```ini
[baidu]
appid = xxxxxxxxxxxxxxxxx
key = xxxxxxxxxxxxxxxxxxxx
;源语言 默认auto
from = jp
;目标语言
to = zh
[mysql]
;是否开启数据库缓存,避免同一个词被多次查询
;switch = on
;数据库用户名
;user = zen
;数据库密码
;passwd = 163453
;要使用的数据库
;database = mydb
;数据库地址
ip = 192.168.1.5
;数据库端口
;port = 3306
;apk add translate-shell
;apt-get install translate-shell
;dnf install translate-shell
;brew install translate-shell
[shell]
from = ja
to = zh
proxy = 192.168.1.20:8889
```

# 改用sqllite作为数据库

文件名`trans.db`