package vaccineDemographics

type Area struct {
	State       string
	Name4       string
	CensusCode  int
	CensusStats AreaCensusStats
}

func NewArea(state string, name4 string) *Area {
	return &Area{
		State: state,
		Name4: name4,
		CensusStats: AreaCensusStats{
			Religion: *NewReligion(),
		},
	}
}

type AreaCensusStats struct {
	Population                      int
	Age                             Age
	NumFamiles                      int
	Income                          Incomes
	AvgPeoplePerHousehold           float64
	MedianWeeklyRent                float64
	AverageMotorVehiclesPerDwelling float64
	// Education                   Education
	Religion     Religion
	HoursWorked  HoursWorked
	TravelToWork TravelToWork
}

type Religion struct {
	Raw map[string]int
}

func NewReligion() *Religion {
	return &Religion{Raw: make(map[string]int)}
}

func (A *Area) GetReligionPercentages() map[string]float64 {
	percentageMap := make(map[string]float64)
	for religion, numPeople := range A.CensusStats.Religion.Raw {
		percentageMap[religion] = float64(numPeople) / float64(A.CensusStats.Population)
	}
	return percentageMap
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
