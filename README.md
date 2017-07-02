# R
读取trello中的数据生成周报

## 配置文件~/.R/config.ini
```
[authorize]
; trello的appkey
appKey = APPKEY
; trello的token
token = TOKEN

[filter]
; trello中用户名
username = trello

[path]
; 生成markdown文件的地址
output = /path/to/

[mail]
; 发件人
from = user1@example.com
; 收件人
to = user2@example.com
; 密码
pwd = PASSWORD
; 别名
alias =
; 服务器地址
host = smtp.example.com
; 服务器端口
port 587
; 邮件主题
subject =
; 邮件内容
body =
```
