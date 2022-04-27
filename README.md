# bookmark
## gin + vue 实现的 web书签管理器
### 一个web书签管理系统，后端使用gin框架，redis作缓存，前端使用vue。支持书签分类的任意多级创建，书签的批量添加，移动等，实现创建分类自动搭配背景图，配合chrome插件可以一键收藏网页

#### 项目部署地址  [登录/注册](http://asilentboy.cn/bookmark/#/login)          [个人空间](http://asilentboy.cn/bookmark/#/space)

#### 新用户注册

![](http://raxadls58.hn-bkt.clouddn.com/register.jpg)

#### 个人空间

![](http://raxadls58.hn-bkt.clouddn.com/space.jpg)



#### 右键菜单一键收藏

![](http://raxadls58.hn-bkt.clouddn.com/右键菜单.jpg)



#### 通过浏览器插件可以进入个人空间

![](http://raxadls58.hn-bkt.clouddn.com/plugin.jpg)



## 说明

#### 服务器配置文件

```yaml
server:
  port: 9000
  mode: debug

#mysql配置
datasource:    
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: xxx    
  username: xxx
  password: xxx

#redis配置
redis:         
  host: 127.0.0.1
  port: 6379
  password: xxx
  dbName: 0

#用于生成jwt token的secret key
jwt:           
  secretKey: xxx

#用于发送注册验证码的邮箱服务器配置(我自己部署中使用的qq邮箱)
email:         
  username: xxx
  sender: xxx
  password: xxx
  host: xxx

#对象存储配置(我自己部署中使用的七牛云)
oss:          
  server: xxxxxx
  bucket: xxx
  accessKey: xxxxxxxxxxxxxxxxxxxx
  secretKey: xxxxxxxxxxxxxxxxxxxx
```



#### 因为系统会在用户创建书签文件夹(分类)的时候根据文件夹名自动进行配图，所以添加了一个专门的配置文件

```yaml
# 没有匹配到预定义的关键词，系统在以下图源中获取一张风景图作为配图(这个暂时未启用)
pictureSource:
  source1:
    urlPattern: https://www.colorhub.me/search?tag={search}&page={page}
    detailPattern: //div[@class='card photo-card']/a
    targetPattern: //img[@class='card-img-top']

# 下面是一些预定义关键词(这里只添加了一小部分)，创建文件夹名包含某个关键词时，系统随机挑选一张对应关键词下的配图
keyword:
  算法:
    - category_algorithm_bg1.jpg  #在云存储中的key
    - category_algorithm_bg2.png
    - category_algorithm_bg3.jpg

  c++:
    - category_cpp_bg1.jpg
    - category_cpp_bg2.jpg
    - category_cpp_bg3.jpg
    - category_cpp_bg4.jpg

  golang:
    - category_golang_bg1.png
    - category_golang_bg2.jpg
    - category_golang_bg3.jpg

  java:
    - category_java_bg1.jpg
    - category_java_bg2.jpg
    - category_java_bg3.jpg
    - category_java_bg4.jpg

  linux:
    - category_linux_bg1.jpg
    - category_linux_bg2.jpg
    - category_linux_bg3.jpg

  ubuntu:
    - category_ubuntu_bg1.jpg
    - category_ubuntu_bg2.jpg

  mysql:
    - category_mysql_bg1.jpg
    - category_mysql_bg2.jpg
```


### 本来是自用的，结果写成了带用户注册的。目前前端部分还有小部分功能，比如说查看回收站，恢复回收站等没有实现

