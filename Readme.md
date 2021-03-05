# docker remote api 后端服务

## 该服务go语言开发，目前该服务在开发过程中

启动服务, 针对Linux或者MacOS
```bash
cd dockerui && go build
./godocker
```

## 接口调用

### 1、查询镜像接口

**查询所有镜像**

```html
http://ip:28080/images/all
```

**分页**
```html
http://ip:28080/images/all?skip=0&limit=10
```

**查询某一镜像**
```html
http://ip:28080/images/search?image=mysql
```

### 2.查询容器接口

**查询所有容器**

```html
http://ip:28080/container/all
```

**分页**

```html
http://ip:28080/container/all?skip=0&limit=10
```

**查询退出的容器**

```html
http://ip:28080/container/search?quiet=exited&skip=0&limit=10
```

**查询某一容器**

```html
http://ip:28080/container/search?container=redis
```

**感兴趣的小伙伴可以加入，目前想寻前端小伙伴一起完成此项目**