package go_lru

import (
    "fmt"
    "strconv"
    "testing"
    "time"
)

func TestCache(t *testing.T) {

    l := Lru{MaxSize: 10, Cache: make(map[interface{}]interface{}), OldCache: make(map[interface{}]interface{})}

    for i := 0; i < 20; i++ {
        l.Set("hello"+strconv.Itoa(i), "world"+strconv.Itoa(i))
    }

    time.Sleep(time.Second)

    l.Set("hello", "world")

    fmt.Println(l.Cache)
    fmt.Println(l.Get("hello"))
    fmt.Println(l.Get("hello10"))

    fmt.Println(l.Has("hello15"))
}
