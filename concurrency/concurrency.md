# 并发编程

[TOC]

Author: Liu Lixiang
Date: 2019-07-26

## 1. 协程机制

Thread VS Goroutine:

1. 创建时默认stack大小：
    - JDK5以后Java Thread Stack大小为1m
    - Goroutine的stack初始化大小为2k
2. 与KSE(Kernel Space Entity)对应关系
    - Java Thread是1：1
    - Goroutine是M：N
    
## 2. 共享内存机制

传统方式，后续使用csp，但是这个也要了解

```go
sync.Mutex
sync.RWLock
```

类似Java
```java
Lock lock = ...;
lock.lock();
try{
    //thread safe
} catch (Exception ex) {

} finally {
    lock.unlock();
}

```

## 3. CSP并发机制

