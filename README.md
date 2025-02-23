Talim
======

twitter data copier

周五的时候为了抓公司竞品数据，才知道油猴脚本里劫持 xhr 非常简单，于是打算做个之前要搞的东西：备份 twitter 数据。

周末搞了两天搞了个虽然简单但已经能用的东西，最终抓到的数据可以看 <https://talim.9farm.com>

抓取脚本在 [public/twitter-copier.user.js](/public/twitter-copier.user.js)，我只会用 [Tampermonkey](https://chromewebstore.google.com/detail/tampermonkey/dhdgffkkebhmkfjojejmpbldmpobfkfo)

本来我想的是大家装了同一个脚本后往同一个数据写入，这样交叉打数据会更全，但马上意识到有严重安全问题，可能保存的数据里有隐私的东西，包括打数据的时间本身都是，所以这个脚本只供演示和研究用途。

我自己有一套复杂的 golang 习惯，不太确认别人是否接受

简易的安装过程是，如果只在 docker 里尝试，需要自己有 https 证书

1. 需要 mysql，在 `/misc/db` 目录里有 `user-docker.sql` 创建用户权限，如果不同机器要修改允许 ip，`talim-struct.sql` 建表
2. `misc/docke` 目录有 `Dockerfile`，打包方法类似 `sudo docker build -t talim -f Dockerfile ../..`
3. `/misc/docker/pull.sh` 是我现在正在用的脚本，没空改了仅作参考
3. 网页客户端用的 angular，可有可无，只是演示用，我关心的是数据。在 `/client` 目录 `npm i` 安装依赖后，先在 `/proto` 目录执行 `make client` 来生成必要的 pb 文件，然后回到 `/client` 目录敲 `make` 会生成 `/client/config.ini` 文件，可以修改后重新 `make` 启动客户端
4. nginx 配置在 `/misc/nginx/dev.conf`，需要修改对应的证书文件、路径、端口

服务器端非常简单，属于不设防，也没做什么性能优化，希望不要搞 flood 破坏，当然如果不是 DoS 而是能黑进去，还是非常希望能指教一二的。

------

因为 x.com 现在没有公开 api，没空详细解数据结构屎山，有两点估计是可行的，但不知道需要多少精力/时间：

1. 可能能分析返回数据，直接从源头删除广告
2. 搞明白 graphql 参数后直接调想要的数据，我现在只能保存所有浏览器看到的内容

我现在是在 chrome console 里输自动滚屏来抓数据，但滚不了几十下就 429 了

```
setInterval(() => {
  window.scrollTo(0, document.body.scrollHeight);
}, 5000);
```
