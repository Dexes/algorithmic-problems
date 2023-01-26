package main

import (
	"math/rand"
	"time"
)

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

type Key [3]byte

func makeKey() Key {
	return Key{makeLetter(), makeLetter(), makeLetter()}
}

func toKey(s string) Key {
	return Key{s[0], s[1], s[2]}
}

func (x Key) String() string {
	return string([]byte{x[0], x[1], x[2]})
}

func makeLetter() byte {
	return byte(random.Intn(26) + 'a')
}
