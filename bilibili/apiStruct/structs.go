package apiStruct

// PublishedListResponse B站已上架<商品列表>响应数据结构体
type PublishedListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Total             int           `json:"total"`
		List              []PublishItem `json:"list"`
		PageNum           int           `json:"pageNum"`
		PageSize          int           `json:"pageSize"`
		Size              int           `json:"size"`
		StartRow          int           `json:"startRow"`
		EndRow            int           `json:"endRow"`
		Pages             int           `json:"pages"`
		PrePage           int           `json:"prePage"`
		NextPage          int           `json:"nextPage"`
		IsFirstPage       bool          `json:"isFirstPage"`
		IsLastPage        bool          `json:"isLastPage"`
		HasPreviousPage   bool          `json:"hasPreviousPage"`
		HasNextPage       bool          `json:"hasNextPage"`
		NavigatePages     int           `json:"navigatePages"`
		NavigatepageNums  []int         `json:"navigatepageNums"`
		NavigateFirstPage int           `json:"navigateFirstPage"`
		NavigateLastPage  int           `json:"navigateLastPage"`
	} `json:"data"`
	Errtag int `json:"errtag"`
}

// PublishItem B站已上架<商品>响应数据结构体
type PublishItem struct {
	C2CItemsId    int64  `json:"c2cItemsId"`
	Type          int    `json:"type"`
	C2CItemsName  string `json:"c2cItemsName"`
	DetailDtoList []struct {
		BlindBoxId  int    `json:"blindBoxId"`
		ItemsId     int    `json:"itemsId"`
		SkuId       int    `json:"skuId"`
		Name        string `json:"name"`
		Img         string `json:"img"`
		MarketPrice int    `json:"marketPrice"`
		Type        int    `json:"type"`
		IsHidden    bool   `json:"isHidden"`
	} `json:"detailDtoList"`
	TotalItemsCount int         `json:"totalItemsCount"`
	Price           float32     `json:"price"`
	ShowPrice       string      `json:"showPrice"`
	ShowMarketPrice string      `json:"showMarketPrice"`
	Uid             string      `json:"uid"`
	PaymentTime     int         `json:"paymentTime"`
	IsMyPublish     bool        `json:"isMyPublish"`
	Uface           interface{} `json:"uface"`
	Uname           interface{} `json:"uname"`
	UspaceJumpUrl   string      `json:"uspaceJumpUrl"`
}

// DropItemResponse B站下架操作响应数据结构体
type DropItemResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Success            bool        `json:"success"`
		InValidBlindBoxIds interface{} `json:"inValidBlindBoxIds"`
		ErrMsg             string      `json:"errMsg"`
	} `json:"data"`
	Errtag int `json:"errtag"`
}

// CheckResponse B站上架前检测响应数据结构体
type CheckResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status             bool        `json:"status"`
		InvalidMsg         string      `json:"invalidMsg"`
		ShowTime           string      `json:"showTime"`
		ExpectTime         int         `json:"expectTime"`
		Token              string      `json:"token"`
		InvalidBlindBoxIds interface{} `json:"invalidBlindBoxIds"`
	} `json:"data"`
	Errtag int `json:"errtag"`
}

// PublishResponse B站上架操作响应数据结构体
type PublishResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status             string      `json:"status"`
		C2CId              int64       `json:"c2cId"`
		InValidBlindBoxIds interface{} `json:"inValidBlindBoxIds"`
		ErrMsg             string      `json:"errMsg"`
		ConfirmInfo        struct {
			PublishText string `json:"publishText"`
			WaitTime    int    `json:"waitTime"`
			Price       string `json:"price"`
			Discount    string `json:"discount"`
			GoodsCount  int    `json:"goodsCount"`
		} `json:"confirmInfo"`
	} `json:"data"`
	Errtag int `json:"errtag"`
}
