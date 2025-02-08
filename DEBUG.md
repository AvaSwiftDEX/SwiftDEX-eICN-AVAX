1. 使用 geth dev mode 时, 交易的 gasPrice 要大于 10, 如果太小的话, geth 会进入死循环状态
2. geth 的版本使用的是 v1.14.12 或者 v1.15.0
3. ganache 也支持事件监听
4. 使用事件监听时, 有两种函数, 一种是 FilterLogs, 一种是 SubscribeFilterLogs, 前一种返回的是一个数组, 而第二种返回的是一个订阅器
