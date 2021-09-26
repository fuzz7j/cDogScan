## cDogScan 

多服务口令爆破、内网常见服务未授权访问探测，端口扫描

### 功能

- 端口扫描
- 暴力破解（ftp、ssh、smb、mongodb、mssql、mysql、postgresql）
- 未授权访问（zookeeper、elastic、memcached、couchdb）

### usage

```
爆破 10.211.55.1-255 ftp、mssql服务
./cDogScan -i 10.211.55.1-255 -m ftp,mssql

使用指定用户名密码字典爆破 10.211.55.1-255 全部服务，不生成result.txt
./cDogScan -i 10.211.55.1-255 -userfile user.txt -passfile pass.txt -nooutput

指定IP文件扫描全部服务
./cDogScan -f ip.txt

仅扫描端口
./cDogScan -i 10.211.55.1 -p 1-65535 -no
```

完整参数

```
  -f string
        ip file, for example: -f ip.txt
  -i string
        ip address,for example: 192.168.11.11 | 192.168.11.11-255
  -m string
        scan type ,for example: -m ssh | -m ssh,ftp,mysql (default "all")
  -no
        disable models, just scan port
  -nooutput
        not output result
  -p string
        port,for example: 22 | 1-65535 (default "21,22,445,1433,2181,3306,5432,5984,6379,9200,27017")
  -pass string
        password
  -passfile string
        password dict, for example: -passfile pass.txt
  -t int
        thread (default 600)
  -time int
        timeout (default 3)
  -user string
        username
  -userfile string
        username dict, for example: -userfile user.txt
```