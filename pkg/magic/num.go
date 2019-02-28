package magic

import (
    "log"
    "math/rand"
    "strconv"
    "time"
)

type num struct {
    seed *rand.Rand
}

func (n *num) RandInt(min int, max int) int {
    if n.seed == nil {
        n.seed = rand.New(rand.NewSource(time.Now().UnixNano()))
    }
    return n.seed.Intn(max-min) + min
}

func (*num) RandIntArr(min int, max int, length int) []int {
    if max-min < length {
        log.Println("RandIntArr : infinite loop warning! enum less than length")
        return []int{}
    }
    arr := []int{}
    for i := 0; i < length; i++ {

        arr = Arr.AppendIntIfMissing(arr, Num.RandInt(min, max))
        if len(arr) == i {
            // 当前数字已存在，则放弃此轮计数
            i--
        }

    }
    return arr
}

func (n *num) RandIntArrButExcept(min int, max int, length int, excepts []int, arr []int) []int {
    if max-min-len(excepts) < length {
        log.Println("RandIntArrButExcept : infinite loop warning! enum less than length")
        return []int{}
    }

    rands := n.RandIntArr(min, max, length)
    for _, r := range rands {

        if !Arr.HasInt(r, excepts) {
            arr = Arr.AppendIntIfMissing(arr, r)
            if len(arr) == length {
                break
            }
        }
    }

    if len(arr) < length {
        return n.RandIntArrButExcept(min, max, length, excepts, arr)
    }

    return arr
}

func (*num) Uint64ArrToIntArr(arr []uint64) []int {
    r := []int{}
    for _, a := range arr {
        r = append(r, int(a))
    }
    return r
}

func (*num) IntArrToUint64(arr []int) []uint64 {
    r := []uint64{}
    for _, a := range arr {
        r = append(r, uint64(a))
    }
    return r
}

func (*num) IntToString(n int) string {
    return strconv.Itoa(n)
}

func (n *num) U64ToString(target uint64) string {
    return n.IntToString(int(target))
}
