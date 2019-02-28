package magic

type arr struct {
}

func (*arr) AppendIntIfMissing(slice []int, i int) []int {
    for _, ele := range slice {
        if ele == i {
            return slice
        }
    }
    return append(slice, i)
}

func (*arr) HasInt(target int, arr []int) bool {
    for _, v := range arr {
        if v == target {
            return true
        }
    }
    return false
}

func (*arr) HasUint64(target uint64, arr []uint64) bool {
    for _, v := range arr {
        if v == target {
            return true
        }
    }
    return false
}

func (*arr) HasString(target string, arr []string) bool {
    for _, v := range arr {
        if v == target {
            return true
        }
    }
    return false
}
