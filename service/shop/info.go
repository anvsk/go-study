package shop

type info struct {
    UserName string
    Age      int
    banlance int //余额
}

func Info() info {
    return info{
        "anvsky",
        29,
        100000,
    }
}
