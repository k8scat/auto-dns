# rpi-auto-dns

树莓派自动解析 DNS（阿里云域名）

## 启动脚本

修改 `/etc/rc.local`，添加以下内容：

```bash
#!/bin/bash
set -e

ACCESS_KEY="<Aliyun Access Key>"
ACCESS_SECRET="<Aliyun Access Secret>"

DOMAIN="<域名>"
RR="<主机记录>"

for i in {1..60}; do
  IP="$(hostname -I | awk '{print $1}')" # 多个网络
  if [[ -n "${IP}" ]]; then
    break
  fi
  sleep 1
done

/usr/bin/rpi-auto-dns \
    -access-key "${ACCESS_KEY}" \
    -access-secret "${ACCESS_SECRET}" \
    -domain "${DOMAIN}" \
    -rr "${RR}" \
    -ip "${IP}" >/var/log/rpi-auto-dns.log 2>&1
```

> 如果 `/etc/rc.local` 文件不存在，创建后需要添加可执行权限：`chmod a+x /etc/rc.local`。
