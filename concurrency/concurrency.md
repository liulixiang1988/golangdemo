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

类似Java:

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

Java中的Future：

```java
import java.util.concurrent.ExecutionException;
import java.util.concurrent.FutureTask;

public class Main {

    private static FutureTask<String> service() {
        FutureTask<String> futureTask = new FutureTask<>(() -> "Do something");
        new Thread(futureTask).start();
        return futureTask;
    }

    public static void main(String[] args) throws ExecutionException, InterruptedException {
        FutureTask<String> task = service();
        System.out.println("hello");
        System.out.println(task.get());
    }
}
```

## 4. 多路选择与超时

gocurr2_test.go

## 5. channel关闭和广播

gocurr3_test.go

在多个函数内传递sync.WaitGroup时需要用指针，否则会引发deadlock panic

一般由数据生产者关闭channel

`close(chan)`会被所有数据接受者接收

## 6. 任务取消

gocurr4_test.go