package main



/*
    TYPES
 */

type Counter interface {
	Add(uint64)
	Read() uint64
}

type FloatCounter interface {
	Add(float64)
	Read() float64
}



type NotSafeCounter struct {
    value uint64
}

func NewNotSafeCounter() Counter {
    return &NotSafeCounter{0}
}

func (c *NotSafeCounter) Add(num uint64) {
    c.value = c.value + num
}

func (c *NotSafeCounter) Read() uint64 {
    return c.value
}
