package main

import (
	"math/rand"
	"time"
)

var lettersRange = []int{'a', 'z' + 1 - 'a'}
var random = rand.New(rand.NewSource(time.Now().UnixMilli()))

type Codec struct {
	storage map[Key]string
}

func Constructor() Codec {
	return Codec{storage: make(map[Key]string)}
}

func (x *Codec) encode(longUrl string) string {
	for {
		key := makeKey()
		if _, ok := x.storage[key]; ok {
			continue
		}

		x.storage[key] = longUrl
		return key.String()
	}
}

func (x *Codec) decode(shortUrl string) string {
	key := toKey(shortUrl)
	result := x.storage[key]
	delete(x.storage, key)

	return result
}

type Key struct {
	a, b, c byte
}

func makeKey() Key {
	return Key{makeLetter(), makeLetter(), makeLetter()}
}

func toKey(s string) Key {
	return Key{s[0], s[1], s[2]}
}

func makeLetter() byte {
	return byte(random.Intn(lettersRange[1]) - lettersRange[0])
}

func (x Key) String() string {
	return string([]byte{x.a, x.b, x.c})
}
