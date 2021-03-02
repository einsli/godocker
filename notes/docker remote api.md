# docker remote api 开发

##  一、配置开启docker remote api

### 1. 开启SSL证书远程加密连接

> 备注: docker ssl 证书参考下面链接，中间也做了一些修改
>
> docker 开启api可以开启ssl验证，但是为了安全，开发过程中也需要指定证书，因此建议配置的时候使用证书

**来源:** <a href="https://blog.csdn.net/xu_cxiang/article/details/104529712">Docker 配置SSL证书加密远程链接 Remote/Rest API</a>

**1.1** 在服务端 /etc目录下创建docker 目录，并切换到该目录下

```bash
mkdir /etc/docker && cd /etc/docker
```

**1.2** 创建根证书RSA私钥

```bash
openssl genrsa -aes256 -out ca-key.pem 4096
```

> 备注：此处会提示输入证书密码，一共需要输入两次，请根据实际情况设置密码，设置成功后目录下生成ca-key.pem密钥文件

**1.3**创建CA证书

```bash
openssl req -new -x509 -days 1000 -key ca-key.pem -sha256 -subj "/CN=*" -out ca.pem
```

> 备注：以上一步生成的ca-key.pem秘钥创建证书，这里是自己作为ca机构，自己给自己签发证书，也可以从第三方ca机构服务商处签发证书。

**1.4** 创建服务端私钥

```bash
openssl genrsa -out server-key.pem 409
```

> 备注：此处生成server-key.pem密钥（服务端私钥）。

**1.5** 创建服务端签名请求证书文件

```bash
openssl req -subj "/CN=*" -sha256 -new -key server-key.pem -out server.csr
```

> 备注：此处生成服务端证书文件server.csr 。

**1.6** 创建签名生效的服务端证书文件

```bash
openssl x509 -req -days 1000 -sha256 -in server.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem
```

> 备注：此处签名后的正式server-cert.pem为签名生效的服务端证书。创建期间要求输入证书密码（之前创建的证书密码）

**1.7** 创建客户端私钥

```bash
openssl genrsa -out key.pem 4096
```

> 备注：此处生成的key.pem文件为客户端私钥，用于客户端远程链接认证

**1.8**

```bash
openssl req -subj "/CN=client" -new -key key.pem -out client.csr
```

> 备注：此处生成的为客户端签的证书文件client.csr。

**1.9** 创建extfile.cnf的配置文件

```bash
echo extendedKeyUsage=clientAuth > extfile.cnf
```

**1.10 **创建签名生效的客户端证书文件

```bash
openssl x509 -req -days 1000 -sha256 -in client.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out cert.pem -extfile extfile.cnf
```

> 备注：此处生成的为客户端证书文件，用于客户端远程链接认证。

**1.10** 删除多余文件

```bash
rm -rf ca.srl client.csr extfile.cnf server.csr
```

> 备注：删除多余文件后，该目录下剩余：
> ca.pem CA机构证书
> ca-key.pem 根证书RSA私钥
> cert.pem 客户端证书
> key.pem 客户私钥
> server-cert.pem 服务端证书
> server-key.pem 服务端私钥

### 2. 配置Docker支持TSL链接

**2.1** 编辑docker.service配置文件

```bash
vim /lib/systemd/system/docker.service
```

> 备注: 上述文件路径是针对centos系统，若操作系统为Ubuntu请执行以下命令

```bash
vim /etc/systemd/system/docker.service
```

**找到ExecStart = 开头的一行代码，将其替换为如下内容：**

```
ExecStart=/usr/bin/dockerd -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock --tlsverify --tlscacert=/etc/docker/ca.pem \
 10 --tlscert=/etc/docker/server-cert.pem --tlskey=/etc/docker/server-key.pem
```

> 备注：此处指定了ca证书、服务端证书和服务端密钥，端口设置为：2375（docker默认端口）

**2.2** 刷新配置，重启docker

```bash
systemctl daemon-reload && systemctl restart docker
```

### 3. 验证TSL方式远程链接Docker

**3.1** 将服务器/etc/docker目录下的ca.pem、cert.pem、key.pem三个文件复制到客户端

下面以macOS操作系统为例

```bash
mkdir ~/.docker && cd ~/.docker 
scp root@serverip:/etc/docker/{ca,cert,key}.pem ./
```

**3.2 ** 在终端远程测试

```bash
cd ~/.docker
curl ip:2379/images/json  --cert ./cert.pem --key ./key.pem -k
```

> 备注: docker 查看镜像接口  ip:port/images/json 
>
> 因为我们配置了秘钥，所以需要携带证书访问
>
> 如果输出镜像信息则说明配置成功

如果想要以json格式化输出，可以执行以下命令

```bash
cd ~/.docker
curl ip:2379/images/json  --cert ./cert.pem --key ./key.pem -k |  python -mjson.tool
```



**至此 docker remote api 远程访问已经配置完成，后续会更新api开发章节**