package vaccineDemographics

type Area struct {
	State       string
	Name4       string
	CensusCode  int
	CensusStats AreaCensusStats
}

type AreaCensusStats struct {
	Population                  int
	Age                         Age
	NumFamiles                  int
	Income                      Incomes
	AvgPeoplePerHousehold       float32
	MedianWeeklyRent            float64
	AvgMotorVehiclesPerDwelling float32
	// Education                   Education
	Religion     map[string]float32
	HoursWorked  HoursWorked
	TravelToWork TravelToWork
}

type HoursWorked struct {
	Pct1to15Hours  float32
	Pct16to24Hours float32
	Pct25to34Hours float32
	Pct35to39Hours float32
	Pct40PlusHours float32
}

type Incomes struct {
	MedianPersonal  float64
	MedianFamily    float64
	MedianHousehold float64
}

type TravelToWork struct {
	CarDriver       float32
	CarPassenger    float32
	WorkAtHome      float32
	Bus             float32
	Train           float32
	PublicTransport float32
}

type Age struct {
	Median    int
	Num0to4   int
	Num5to9   int
	Num5to14  int
	Num15to19 int
	Num20to24 int
	Num25to34 int
	Num35to44 int
	Num45to54 int
	Num55to64 int
	Num65to74 int
	Num75to84 int
	Num85Plus int
}
