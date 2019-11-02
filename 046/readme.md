# 知识点

IPC (Inter-Process Communication)，表示进程间通信的方法。主要包括有：**系统信号量**(signal)，**管道**(pipe)，**套接字**(socket)，**文件锁**(file lock)，**消息队列**(message queue)，**信号灯**(semaphore)。

Go 中的 net 包，大量地、间接或直接使用到 syscall.Socket 函数。

1. net.Dial 函数的第一个参数 network 有哪些可选值？

network 的可选值为 string 类型，代表了底层创建 socket 用到的协议。
- tcp：tcp 协议，基于的 ip 协议版本根据 address 的值自适应。
- tcp4：基于 ip 协议第四版本的 tcp。
- tcp6：基于 ... 第六版 ... 。
- udp：udp 协议，基于的 ip 协议版本根据 address 的值自适应。
- udp4：基于 ip 协议第四版本的 udp。
- udp6：基于 ... 第六版 ... 。
- unix：unix 通信域下的内部 socket 协议。以 SOCK_STREAM 为 socket 类型。
- unixgram：... 。以 SOCK_DGRAM 为 socket 类型。
- unixpacket：...。以 SOCK_SEQPACKET 为 socket 类型。

syscall.Socket() 接收三个参数：
- domain(通信域)：AF_INET、AF_INET6、AF_UNIX。
- typ(类型)：SOCK_DGRAM、SOCK_STREAM、SOCK_SEQPACKET、SOCK_RAM。
- proto(协议)

SOCK_DGRAM：有消息边界，但没有逻辑连接的非可靠 socket 类型。
- 有消息边界：内核发送或接收数据的时候以消息为单位(内核负责控制)。

SOCK_STREAM：没有消息边界，有逻辑连接。能够保证传输的**可靠性**和数据的**有序性**。还能实现数据的**双向传输**。
- 字节流形式传输数据，以字节为单位。内核不知道一段字节流中有多少消息，也不知道消息是否完整，需要应用程序自己控制。

![socket](/046/png/99f8a0405a98ea16495364be352fe969.png)

2. 调用 net.DialTimeout 函数给定的超时意味着什么？

**函数为网络连接建立完成而等待的最长时间**。

如果有必要，会访问 DNS 服务。如果解析出来的是多个 IP，函数会串行或并发的尝试连接，函数总以最先建立连接的那个连接为准。

同时，它还会根据超时前的剩余时间，去设定针对每次连接尝试的超时时间，以便让它们都有适当的时间执行。

# 服务器与客户端通行(TCP)  (code: server.go)

# 访问百度(HTTP)  (code: connbaidu.go)
