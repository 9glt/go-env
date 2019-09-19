package env

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

var errNotFound = errors.New("ENV Variable not found")

func Get(k string) *Item {
	s, exists := os.LookupEnv(k)
	if !exists {
		log.Fatal(errNotFound)
	}
	return &Item{s}
}

type Item struct{ v string }

func (i *Item) Int() int {
	n, err := strconv.Atoi(i.v)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func (i *Item) Float() float64 {
	f, err := strconv.ParseFloat(i.v, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func (i *Item) String() string {
	return i.v
}

func (i *Item) Duration() time.Duration {
	d, err := time.ParseDuration(i.v)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
