package shop

type goods struct {
    Name   string `json:"name"`
    Stocks int    `json:"stocks"`
    Price  int    `json:"price"`
}

type ReqList struct {
    Page  int `form:"page"`
    Limit int `form:"limit"`
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
        req.Limit = 10
    }
    if req.Limit > len(g) {
        req.Limit = len(g)
    }

    return g[:req.Limit]
}
