# Unifi 工具箱

[English](./README.md) | **中文**

在Unifi OS和控制器的网页控制台中，有部分功能是缺失的，导致使用上不方便

本项目提供了以下工具，以提升控制台的使用体验：

### 1. 清理离线客户端
有很多访客客户端残留在客户端列表，一个个删除太麻烦。<br>
此工具可以批量删除，只有那些没有被自定义命名的客户端会被清理。
工具的删除操作跟在页面上点"Remove“是一样的。

### 2. 打印mac过滤地址列表
Mac地址过滤列表中，只展示了mac地址而没有客户端名称，辨别起来不方便<br>
此工具可以打印所有WiFi下的mac过滤列表，包括mac地址和客户端名称

## 参数
- -m: 设备类型, 必填<br>
  只能填**Console** 或者 **Controller**<br>
  Console 指类似 UDMPro/UDR等设备<br>
  Controller 指自部署的controller
- -g: Console/Controller IP 地址, **必填**
- -p: Console/Controller 端口, **选填**. 默认值是 **443**
- -u: Console/Controller 用户名, **必填**
- -d: Dry Run, **选填**. 默认值是 **false**

## 例子
```shell
./print-mac-filter -m Console -g 192.168.10.1 -u xxx

./prune-clients -m Console -g 192.168.10.1 -u xxx -d true # 清理前运行
./prune-clients -m Console -g 192.168.10.1 -u xxx
```
