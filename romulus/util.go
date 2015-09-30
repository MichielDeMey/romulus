package romulus

import (
	"math/rand"
	"strings"
	"time"
)

var unicodeReplacements = map[string]string{
	`\u003c`: "<",
	`\u003e`: ">",
	`\u0026`: "&",
}

// HTMLUnescape reverses the HTMLEscape process done by JSON encoding
func HTMLUnescape(s string) string {
	r := s
	for k, v := range unicodeReplacements {
		r = strings.Replace(r, k, v, -1)
	}
	return r
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	charBits = 6
	charMask = 1<<charBits - 1
	charMax  = 63 / charBits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStr(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), charMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), charMax
		}
		if idx := int(cache & charMask); idx < len(chars) {
			b[i] = chars[idx]
			i--
		}
		cache >>= charBits
		remain--
	}

	return string(b)
}
