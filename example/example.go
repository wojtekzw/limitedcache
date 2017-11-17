package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/wojtekzw/limitedcache"
)

func printCacheOperations(c *limitedcache.Cache) {
	ec := c.Events()
	for {
		select {
		case ev := <-ec:
			log.Printf("event: %s, %s, %v", ev.Operation(), ev.Key(), ev.Status())
		}
	}
}

func generateCacheWorkload(c *limitedcache.Cache, limit int) {
	for i := 0; i < limit; i++ {
		c.Set(strconv.Itoa(i), []byte(fmt.Sprintf("content: %d", i)))
	}
}

func countFiles(basePath string) (int, error) {
	var count int
	err := filepath.Walk(basePath, func(path string, f os.FileInfo, err error) error {
		if err == nil && !f.IsDir() {
			count++
		}
		return err
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func main() {
	path := "/tmp/limitedcache-test"
	itemsInCache := 3
	itemsToAdd := 10
	os.RemoveAll(path)
	c := limitedcache.New(path, itemsInCache)
	go printCacheOperations(c)
	generateCacheWorkload(c, itemsToAdd)

	time.Sleep(500 * time.Millisecond)

	count, err := countFiles(path)
	if err != nil {
		log.Printf("error count files: %v", err)
		os.Exit(1)
	}
	log.Printf("number of cached items on disk: %d", count)
}
