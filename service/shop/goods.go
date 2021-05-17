package shop

type goods struct {
    Name   string
    Stocks int
    Price  int
}

type ReqList struct {
    Page  int `form:"page"`
    Limit int `form:"limit"`
    // Type  int `form:"type" binding:"required"`
}

func List(req ReqList) []goods {
    g := []goods{
        {
            "耐克上衣",
            10,
            100,
        },
        {
            "Adidas外套",
            20,
            200,
        },
        {
            "Fila外套",
            10,
            300,
        },
    }
    if req.Limit == 0 {
        req.Limit = 1
    }
    if req.Limit > len(g) {
        req.Limit = len(g)
    }

    return g[:req.Limit]
}
