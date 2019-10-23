# 知识点

**条件变量是基于互斥锁的**，在互斥锁的支持下才能工作。

**条件变量不是用来保护临界区和共享资源的**，而是用来协调想要访问共享资源的那些线程。当共享资源的状态发生变化时，它可以通知被互斥锁锁阻塞的线程。
- 优势：当共享资源不满足条件时，想操作它的线程不用循环的检查状态，等待通知就好。

1. 条件锁与互斥锁如何配合使用？

```go
var lock sync.RWMutex

w_cond = sync.NewCond(&lock)
r_cond = sync.NewCond(lock.RLocker())

...
/* 针对写 */
lock.Lock()
for /* 条件 */ {
    w_cond.Wait()
}
...
lock.Unlock()
r_cond.Signal()


...
/* 针对读 */
lock.RLock()
for /* 条件 */ {
    r_cond.Wait()
}
lock.RUnlock()
w_cond.Signal()
```

条件变量提供的方法有三个：等待通知（wait）、单发通知（signal）和广播通知（broadcast）。

2. 条件变量的 Wait() 方法做了什么？

四件事：
- 将调用它的 goroutine 加入到当前条件变量的通知队列中。
- **解锁当前条件变量基于的那个互斥锁**。
- 让当前 gorutine 处于等待状态，等到通知到来时再决定是否唤醒它。(也就是说，代码阻塞在 Wait 处)
- 当通知到来并决定唤醒 goroutine，则**唤醒之后重新锁定当前条件变量基于的那个互斥锁**。

如果调用 Wait 之前，没有先上锁，则程序引发 panic。如果 Wait 方法不先解锁互斥锁，那么造成两种结果：**程序因 panic 崩溃，或者相关 goroutine 全面阻塞**。

3. 为什么调用 Wait 方法时，要在 for 循环里边。

为保险起见，如果一个 goroutine 因收到通知而被唤醒，但发现共享资源的状态还不符合自己的要求，此时应该再次调用条件变量的 Wait() ，等待下一次唤醒通知。对应实际情况中的：
- 当多个 goroutine 竞争一个共享资源时，且它们的操作逻辑一致，则多个 goroutine 只需要被唤醒一个即可。其他的通过 for 循环再次进入等待队列。
- 当共享资源的状态存在多个时，被唤醒的 goroutine 需要确保当前共享资源的状态就是自己需要的，否则调用 Wait() 重新进入等待队列。
- 在一些多 CPU 核心的计算机系统中，即使没有收到条件变量的通知，调用 Wait 方法的 goroutine 也是有可能被唤醒的。这是由计算机硬件层面决定的，即使是操作系统（比如 Linux）本身提供的条件变量也会如此。

4. 条件变量的 Signal 与 Broadcast 方法的异同点。

Signal 只会唤醒一个 goroutine，Broadcast 会唤醒全部。(条件变量的 Wait 方法总会把当前的 goroutine 添加到通知队列的**队尾**，而它的 Signal 方法总会从通知队列的**队首**开始，查找可被唤醒的 goroutine。)

注意：与 Wait 不同，Signal 和 Broadcast 不需要在互斥锁的保护下执行。从程序的执行效率来看，最好是解锁条件变量基于的互斥锁之后，再去调用 Signal 和 Broadcast。

**条件变量的通知具有即时性。也就是说，如果发送通知的时候没有 goroutine 为此等待，那么该通知就会被直接丢弃。在这之后才开始等待的 goroutine 只可能被后面的通知唤醒。**

优先考虑使用**Broadcast**。