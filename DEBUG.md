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
7. map的key不能用*big.Int, 需要用string, 因为\*big.Int的值虽然相同, 但是地址不同, 所以不能作为key
8. log.Fatal 会直接退出程序, 所以不能随便使用
9. *big.Int 比较大小时, 不能直接用==, 需要用Cmp方法
