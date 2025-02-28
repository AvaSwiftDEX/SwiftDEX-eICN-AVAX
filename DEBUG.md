# Debug Notes

1. 使用 geth dev mode 时, 交易的 gasPrice 要大于 10, 如果太小的话, geth 会进入死循环状态
2. geth 的版本使用的是 v1.14.12 或者 v1.15.0
3. ganache 也支持事件监听
4. 使用事件监听时, 有两种函数, 一种是 FilterLogs, 一种是 SubscribeFilterLogs, 前一种返回的是一个数组, 而第二种返回的是一个订阅器
5. 如果使用 geth v1.10.10, 那么需要安装 golang v1.14 以及 gcc

```bash
sudo apt update
sudo apt install build-essential
gcc --version
```

6. map的key不能用*big.Int, 需要用string, 因为\*big.Int的值虽然相同, 但是地址不同, 所以不能作为key
7. log.Fatal 会直接退出程序, 所以不能随便使用
8. *big.Int 比较大小时, 不能直接用==, 需要用Cmp方法
9. geth v1.15.0 的 basefeePerGas 默认很高，交易如果设置了较低的gasPrice，就无法打包，所以需要设置较低的basefeePerGas，我在 geth 源代码中设置成了 10，baseFeePerGas会在geth运行的过程中根据区块大小不断变化，最低是7. 源代码的位置是 params/protocol_params.go，参数是InitialBaseFee
10. 如果出现`address already in use`错误，使用：

```bash
lsof -i :8080 # 查看端口占用
kill -9 $(lsof -t -i :8080) # 强制释放端口
```

11. 当测试用例异常中断时，可以用这个命令清理：

```bash
ps -ef | grep "go test" | awk '{print $2}' | xargs kill -9
```
