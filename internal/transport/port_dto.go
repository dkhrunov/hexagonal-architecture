package transport

type PortDto struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Code        string     `json:"code"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
	Alias       []string   `json:"alias"`
	Regions     []string   `json:"regions"`
	Coordinates [2]float64 `json:"coordinates"`
	Province    string     `json:"province"`
	Timezone    string     `json:"timezone"`
	Unlocs      []string   `json:"unlocs"`
}
