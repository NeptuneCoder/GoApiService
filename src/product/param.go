package product


type ServiceTypeParam struct {
	ServiceType    string `josn:"serviceType"`
	Ad             string `josn:"ad"`
	Video          string `josn:"video"`
	Speed          string `josn:"speed"`
	Image          string `josn:"image"`
	ServiceExplain string `josn:"serviceExplain"`
	Products       []ProductParam
}


type ProductParam struct {
	serviceType string
	productId   string
	price       string
	priceUnit   string
	validTime   string
	timeUnit    string
}