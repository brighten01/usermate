package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	filePath   = "largefile.bin" // 替换为实际路径
	chunkSize  = 4 * 1024 * 1024 // 4MB 的块大小
	numWorkers = 8               // 根据 CPU 核心数调整
	bufferSize = chunkSize       // 缓冲区大小与块一致
)

type task struct {
	offset int64
	size   int
}

func main() {
	start := time.Now()

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fileSize := fileInfo.Size()
	numChunks := int((fileSize + chunkSize - 1) / chunkSize) // 计算总块数

	pool := &sync.Pool{
		New: func() interface{} {
			buf := make([]byte, bufferSize)
			return &buf
		},
	}

	var wg sync.WaitGroup
	tasks := make(chan task, numChunks)

	// 启动 Worker
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range tasks {
				// 从池中获取缓冲区
				bufPtr := pool.Get().(*[]byte)
				buf := (*bufPtr)[:t.size] // 调整切片长度offset

				// 读取指定块
				_, err := file.ReadAt(buf, t.offset)
				if err != nil {
					panic(fmt.Sprintf("Read error: %v", err))
				}

				// 此处可处理数据（如计算哈希、解析内容等）
				pool.Put(bufPtr) // 放回池中
			}
		}()
	}

	// 分发任务
	for i := 0; i < numChunks; i++ {
		offset := int64(i) * chunkSize
		remaining := fileSize - offset
		size := chunkSize
		if remaining < chunkSize {
			size = int(remaining)
		}
		tasks <- task{offset: offset, size: size}
	}
	close(tasks)

	wg.Wait()
	fmt.Printf("Time taken: %v\n", time.Since(start))
}
