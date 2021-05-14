package ticket

type CommonRes struct {
    Code    int
    Message string
}

type LoginRes struct {
    CommonRes
    Data struct {
        Token  string
        UserID int `json:"userId"`
    }
}

type OrderRes struct {
    CommonRes
    Data struct {
        OrderID string `json:"orderId"`
    }
}

type ListRes struct {
    CommonRes
    Data []ListData
}

type PassengerRes struct {
    CommonRes
    Data []Passenger
}

type SeatClasses struct {
    ClassNum            int         `json:"classNum"`
    ClassName           string      `json:"className"`
    LocalCurrentCount   int         `json:"localCurrentCount"`
    PubCurrentCount     int         `json:"pubCurrentCount"`
    TotalCount          int         `json:"totalCount"`
    FerryPassTotalCount interface{} `json:"ferryPassTotalCount"`
    OriginPrice         float64     `json:"originPrice"`
    TotalPrice          float64     `json:"totalPrice"`
    HalfPrice           float64     `json:"halfPrice"`
    LocalPrice          float64     `json:"localPrice"`
    LocalHalfPrice      float64     `json:"localHalfPrice"`
    SeatState           int         `json:"seatState"`
    SeatStateName       string      `json:"seatStateName"`
    TotalOriginCount    int         `json:"totalOriginCount"`
    CandidateCount      int         `json:"candidateCount"`
}
type ListData struct {
    LineNum          int           `json:"lineNum"`
    LineName         string        `json:"lineName"`
    LineNo           int           `json:"lineNo"`
    ShipName         string        `json:"shipName"`
    StartPortNo      int           `json:"startPortNo"`
    StartPortName    string        `json:"startPortName"`
    EndPortNo        int           `json:"endPortNo"`
    EndPortName      string        `json:"endPortName"`
    SailDate         string        `json:"sailDate"`
    SailTime         string        `json:"sailTime"`
    BusStartTime     string        `json:"busStartTime"`
    Sx               int           `json:"sx"`
    LineDirect       int           `json:"lineDirect"`
    SaleBeginTime    string        `json:"saleBeginTime"`
    SaleEndTime      string        `json:"saleEndTime"`
    StopSaleTime     int           `json:"stopSaleTime"`
    CarStopSaleTime  interface{}   `json:"carStopSaleTime"`
    OnSale           bool          `json:"onSale"`
    OffSaleMsg       interface{}   `json:"offSaleMsg"`
    SeatClasses      []SeatClasses `json:"seatClasses"`
    DriverSeatClass  interface{}   `json:"driverSeatClass"`
    BuyTicketType    int           `json:"buyTicketType"`
    Clxh             int           `json:"clxh"`
    Clxm             string        `json:"clxm"`
    Hxlxh            int           `json:"hxlxh"`
    Hxlxm            string        `json:"hxlxm"`
    Bus              int           `json:"bus"`
    Bus2             int           `json:"bus2"`
    LineState        int           `json:"lineState"`
    LineStateName    string        `json:"lineStateName"`
    EmbarkPortName   string        `json:"embarkPortName"`
    FreeChildCount   int           `json:"freeChildCount"`
    CandidateTimeEnd string        `json:"candidateTimeEnd"`
}

type Passenger struct {
    PassName       string `json:"passName"`
    CredentialType int    `json:"credentialType"`
    ID             int    `json:"id"`
}

type OrderPassenger struct {
    PassName       string `json:"passName"`
    CredentialType int    `json:"credentialType"`
    PassID         int    `json:"passId"`
}

type OrderItemRequests struct {
    OrderPassenger
    SeatClassName  string `json:"seatClassName"`
    SeatClass      int    `json:"seatClass"`
    TicketFee      int    `json:"ticketFee"`
    RealFee        int    `json:"realFee"`
    FreeChildCount int    `json:"freeChildCount"`
}

type Order struct {
    AccountTypeID string `json:"accountTypeId"`
    UserID        int    `json:"userId"`
    BuyTicketType int    `json:"buyTicketType"`
    ContactNum    string `json:"contactNum"`
    ListData
    OrderItemRequests []OrderItemRequests `json:"orderItemRequests"`
    TotalFee          int                 `json:"totalFee"`
    TotalPayFee       int                 `json:"totalPayFee"`
}
