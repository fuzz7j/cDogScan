## cDogScan 

多服务口令爆破、内网常见服务未授权访问探测，端口扫描

### 功能

- 端口扫描
- 口令爆破（ftp、ssh、smb、mongodb、mssql、mysql、postgresql、redis、snmp）
- 未授权访问（zookeeper、elastic、memcached、couchdb）

### usage

```
爆破 10.211.55.1-255 ftp、mssql服务
./cDogScan -i 10.211.55.1-255 -m ftp,mssql

使用指定用户名密码字典爆破 10.211.55.1/24 全部服务，不生成result.txt
./cDogScan -i 10.211.55.1/24 -userfile user.txt -passfile pass.txt -nooutput

指定IP文件扫描全部服务
./cDogScan -f ip.txt

仅扫描端口
./cDogScan -i 10.211.55.1 -p 1-65535 -no
```

完整参数

```
  -T int
        thread (default 600)
  -f string
        ip file, for example: -f ip.txt
  -i string
        ip address,for example: 192.168.0.1 | 192.168.0.1-255 | 192.168.0.1-192.168.255.255 | 192.168.0.1/24
  -m string
        scan type ,for example: -m ssh | -m ssh,ftp,mysql (default "all")
  -no
        disable models, just scan port
  -nooutput
        not output result
  -p string
        port,for example: 22 | 1-65535 (default "21,22,161,445,1433,2181,3306,5432,5984,6379,9200,11211,27017")
  -pass string
        password
  -passfile string
        password dict, for example: -passfile pass.txt
  -t int
        timeout (default 3)
  -user string
        username
  -userfile string
        username dict, for example: -userfile user.txt
```

### Todo

- [x] IP解析格式增加
- [ ] ICMP探活
- [ ] NetBIOS探测、域控探测
- [ ] Web Title获取

### Bug

- 本地编译，扫描MySQL时输出报错信息  
  注释 go/pkg/mod/github.com/go-sql-driver/mysql@v1.6.0/packets.go, Line 37
- SMB不支持2003  
  SMB模块不支持SMBv1，已加入超时跳过

### 更新日志

- 20211122   
  修改端口扫描结果为即时输出  
  增加IP解析格式，-i/-f支持IP输入格式：192.168.0.1 | 192.168.0.1-255 | 192.168.0.1-192.168.255.255 | 192.168.0.1/24
- 20211123  
  默认端口增加161（snmp）  
  增加snmp public弱口令扫描
- 20211125  
  增加FTP匿名账户、MYSQL匿名账户、Redis未授权扫描  
  加入定时输出log及超时跳过，增加进度显示

### Thanks

https://github.com/shadow1ng/fscan  
https://github.com/netxfly/x-crack