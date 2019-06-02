package go_lru

type Lru struct {
    MaxSize  int
    Cache    map[interface{}]interface{}
    OldCache map[interface{}]interface{}
    Size     int
}

func (l *Lru) _set(key string, value interface{}) {
    l.Cache[key] = value
    l.Size++

    if l.Size >= l.MaxSize {
        l.Size = 0
        l.OldCache = l.Cache
        l.Cache = make(map[interface{}]interface{})
    }
}

func (l *Lru) Get(key string) (value interface{}) {
    value = l.Cache[key]
    if value != nil {
        return
    }

    value = l.OldCache[key]
    if value != nil {
        delete(l.OldCache, key)
        l._set(key, value)
        return
    }
    return
}

func (l *Lru) Set(key string, value interface{}) {
    v := l.Cache[key]
    if v != nil {
        l.Cache[key] = value
    } else {
        l._set(key, value)
    }
}

func (l *Lru) Has(key string) bool {
    f1, f2 := false, false
    if l.Cache[key] != nil {
        f1 = true
    }
    if l.OldCache[key] != nil {
        f2 = true
    }

    return f1 || f2
}

func (l *Lru) Delete(key string) {
    v := l.Cache[key]
    if v != nil {
        l.Size--
    }
    delete(l.Cache, key)
    delete(l.OldCache, key)
}

func (l *Lru) Clear() {
    l.Size = 0
    l.Cache = make(map[interface{}]interface{})
    l.OldCache = make(map[interface{}]interface{})
}
