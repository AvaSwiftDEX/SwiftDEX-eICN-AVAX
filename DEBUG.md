1. 使用 geth dev mode 时, 交易的 gasPrice 要大于 10, 如果太小的话, geth 会进入死循环状态
2. geth 的版本使用的是 v1.14.12 或者 v1.15.0
3. ganache 也支持事件监听
4. 使用事件监听时, 有两种函数, 一种是 FilterLogs, 一种是 SubscribeFilterLogs, 前一种返回的是一个数组, 而第二种返回的是一个订阅器
5. 如果使用 geth v1.10.10, 那么需要安装 golang v1.14 以及 gcc
6. 
```bash
sudo apt update
sudo apt install build-essential
gcc --version
```
