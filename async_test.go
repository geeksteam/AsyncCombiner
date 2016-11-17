package async_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/geeksteam/async"
)

func TestString(t *testing.T) {
	c := async.NewCollector(async.StringMap())

	for i := 1; i <= 10; i++ {
		c.Add(func() (async.G, async.G) {
			i := rand.Intn(100)
			time.Sleep(1 * time.Second)
			return fmt.Sprint(i), i
		})
	}

	fmt.Println(c.Result())
}

func TestInt(t *testing.T) {
	c := async.NewCollector(async.IntMap())

	for i := 1; i <= 10; i++ {
		c.Add(func() (async.G, async.G) {
			i := rand.Intn(100)
			time.Sleep(1 * time.Second)
			return i, i
		})
	}

	fmt.Println(c.Result())
}

func TestCustom(t *testing.T) {
	c := async.NewCollector(customMap{})

	for i := 1; i <= 10; i++ {
		c.Add(func() (async.G, async.G) {
			i := rand.Intn(100)
			time.Sleep(1 * time.Second)
			return keyStruct{i}, i
		})
	}

	fmt.Println(c.Result())
}

type keyStruct struct {
	value int
}

type customMap map[keyStruct]interface{}

func (cm customMap) Add(key, value async.G) {
	cm[key.(keyStruct)] = value
}
