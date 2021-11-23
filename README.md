# 七牛云CDN刷新客户端

- 最近频繁使用到七牛云的CDN来做静态资源加速。

- 每一次更新静态资源后，都要去七牛云的CDN控制台上手动点一下刷新缓存。

- 点的次数多了会有点烦，于是就写了一个命令行工具自动化的去刷新CDN。

- 但是二狗同学也想使用，但是他不想自己再新写一个。
  
- 因此就把工具简单的封装了一下，让所有人都可以使用。

# 示例用法：

命令格式：
```bash
cdn_refresh -a <accessKey> -s <secretKey> -u <urlDir>
```

- -a：输入从七牛云获取到的accessKey
- -s：输入从七牛云获取到的secretKey
- -u：输入需要被刷新的URL地址，必须是以/结尾的地址

示例用法：
```bash
./cdn_refresh \
  -a 0UGTdVOlNruYfTopM0DSZk82TAAAAAAAAAAAAAAAA \
  -s 0yHev9o_xkpvDyAAAAAAAAAKQygpnnlamAAAAAAA \
  -u "https://blog.linux.com/"
```

# 推荐的用法：

由于`accessKey`和`secretKey`都很长，我们平时都很难记忆，因此推荐大家把这个工具集成到自己的Shell脚本中。

```bash
#!/usr/bin/env bash



...略



QINIU_ACCESSKEY="0UGTdVOlNruYfTopM0DSZk82TAAAAAAAAAAAAAAAA"
QINIU_SECRETKEY="0yHev9o_xkpvDyAAAAAAAAAKQygpnnlamAAAAAAA"
URLDIR="https://blog.linux.com/"

#调用cdn_refresh工具时，建议使用绝对路径
/opt/cdn_refresh -a ${QINIU_ACCESSKEY} -s ${QINIU_SECRETKEY} -u ${URLDIR}
```