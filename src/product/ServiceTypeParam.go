package product

type ServiceType struct {
	ServiceType    string `josn:"serviceType"`
	Ad             string `josn:"ad"`
	Video          string `josn:"video"`
	Speed          string `josn:"speed"`
	Image          string `josn:"image"`
	ServiceExplain string `josn:"serviceExplain"`
	Products       []Product
}
