package magic

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

var StrSource = rand.NewSource(time.Now().UnixNano())

const (
    LetterBytes   = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    LetterIdxBits = 6                    // 6 bits to represent a letter index
    LetterIdxMask = 1<<LetterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    LetterIdxMax  = 63 / LetterIdxBits   // # of letter indices fitting in 63 bits
)

type str struct {
}

// https://stackoverflow.com/questions/12668681/how-to-get-the-number-of-characters-in-a-string
func (*str) Len(str string) int {
    return len([]rune(str))
}

func (*str) ToUrlParams(params ...string) string {
    r := ""
    // m := map[string]string;
    // for k, v := range params {
    // }

    return r
    // return Contact("&", params...)
}

func (s *str) ToUrl(host string, params string) string {
    return s.Contact("?", host, params)
}

func (*str) Contact(seq string, params ...string) string {
    r := []string{}
    for _, v := range params {
        r = append(r, v)
    }
    return strings.Join(r, seq)
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func (s *str) Random(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for LetterIdxMax characters!
    for i, cache, remain := n-1, StrSource.Int63(), LetterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = StrSource.Int63(), LetterIdxMax
        }
        if idx := int(cache & LetterIdxMask); idx < len(LetterBytes) {
            b[i] = LetterBytes[idx]
            i--
        }
        cache >>= LetterIdxBits
        remain--
    }

    return string(b)
}

func (s *str) RouteReplace(str string, args ...string) string {
    p := strings.Index(str, ":")
    if p < 0 || len(args) == 0 {
        return str
    }
    arr := strings.SplitN(str, ":", 2)
    prefix := arr[0]
    sub := arr[1]

    end := strings.Index(sub, "/")
    if end < 0 {
        end = len(sub)
    }

    r := prefix + args[0] + sub[end:]
    return s.RouteReplace(r, args[1:]...)
}

func (*str) Replace(str string, m map[string]string) string {
    for k, v := range m {
        str = strings.Replace(str, k, v, 1)
    }
    return str
}

func (*str) ToString(n interface{}) string {
    return fmt.Sprint(n)
}

// func (*str) ErrorMsg(topic string, msg string) string {
// 	if !config.Get().IsProduction() {
// 		return Contact(" : ", topic, msg)
// 	}
// 	return topic
// }
