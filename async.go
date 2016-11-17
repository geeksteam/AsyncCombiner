package async

import "sync"

type G interface{}

type Combiner interface {
	Add(key, value G)
}

type result struct {
	key   interface{}
	value interface{}
}

type Collector struct {
	input chan result
	stor  Combiner
	wg    sync.WaitGroup
}

func NewCollector(combiner Combiner) *Collector {
	c := &Collector{
		input: make(chan result, 100),
		stor:  combiner,
	}
	go c.loop()
	return c
}

func (c *Collector) Add(op func() (G, G)) {
	c.wg.Add(1)

	go func() {
		key, value := op()
		c.input <- result{key, value}
	}()
}

func (c *Collector) loop() {
	for r := range c.input {
		c.stor.Add(r.key, r.value)
		c.wg.Done()
	}

}

func (c *Collector) Result() G {
	c.wg.Wait()
	close(c.input)
	return c.stor
}

func StringMap() stringMap {
	return stringMap{}
}

type stringMap map[string]interface{}

func (sm stringMap) Add(key, value G) {
	sm[key.(string)] = value
}

func IntMap() intMap {
	return intMap{}
}

type intMap map[int]interface{}

func (im intMap) Add(key, value G) {
	im[key.(int)] = value
}
