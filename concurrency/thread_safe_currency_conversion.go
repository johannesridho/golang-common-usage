//thread-safe currency conversion

package main

import (
	"fmt"
	"sync"
	"time"
)

var cache *Cache

func main() {
	cache = &Cache{data: make(map[string]float64)}

	c := Currency{CentAmount: 5000, Symbol: "SGD"}
	fmt.Printf("%+v , %.2f \n", c, float64(c.CentAmount)/100)

	c.ConvertTo("IDR")

	fmt.Printf("%+v , %.2f \n", c, float64(c.CentAmount)/100)

	c = Currency{CentAmount: 5000, Symbol: "SGD"}
	fmt.Printf("%+v , %.2f \n", c, float64(c.CentAmount)/100)

	c.ConvertTo("IDR")

	fmt.Printf("%+v , %.2f \n", c, float64(c.CentAmount)/100)
}

type Cache struct {
	data map[string]float64
	mu   sync.RWMutex
}

func (c *Cache) Get(key string) float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *Cache) Set(key string, val float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val
}

type Currency struct {
	CentAmount int
	Symbol     string
}

func (c *Currency) Multiply(e float64) {
	c.CentAmount = int(float64(c.CentAmount) * e)
}

func (c *Currency) Div(e float64) {
	c.CentAmount = int(float64(c.CentAmount) / e)
}

func (c *Currency) String() string {
	return fmt.Sprintf("%d", c.CentAmount/100)
}

func (c *Currency) ConvertTo(symbol string) {
	ch1 := make(chan float64)
	go getUSDConvertionRate(c.Symbol, ch1)

	ch2 := make(chan float64)
	go getUSDConvertionRate(symbol, ch2)

	toUSDRate := <-ch1
	usdToSymbolRate := <-ch2
	usd := int(float64(c.CentAmount) / toUSDRate)
	c.CentAmount = int(float64(usd) * usdToSymbolRate)
	c.Symbol = symbol
}

func getUSDConvertionRate(symbol string, ch chan float64) {
	cached := cache.Get(symbol)
	if cached != 0 {
		fmt.Println("got cached result", cached)
		ch <- cached
		return
	}

	time.Sleep(time.Second * 2)

	var result float64
	if symbol == "IDR" {
		result = 14354.00
	}

	if symbol == "SGD" {
		result = 1.34
	}

	if symbol == "USD" {
		result = 1
	}

	cache.Set(symbol, result)

	ch <- result
}
