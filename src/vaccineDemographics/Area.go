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
			Religion: Religion{Raw: make(map[string]int)},
			Ancestry: Ancestry{Raw: make(map[string]int)},
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
	Ancestry     Ancestry
}

type Ancestry struct {
	Raw map[string]int
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

func (A *Area) GetAncestryPercentages() map[string]float64 {
	percentageMap := make(map[string]float64)
	for key, numPeople := range A.CensusStats.Ancestry.Raw {
		percentageMap[key] = float64(numPeople) / float64(A.CensusStats.Population)
	}
	return percentageMap
}

type HoursWorked struct {
	Average         float64
	Num0            int
	Num1to15        int
	Num16to24       int
	Num25to34       int
	Num35to39       int
	Num40           int
	Num41to48       int
	Num49Plus       int
	NumNotSpecified int
}

func (h *HoursWorked) SetAverage(totalNumPeople int) {
	sumHoursWorked := 0.0
	sumHoursWorked += float64(8 * h.Num1to15)
	sumHoursWorked += float64(20 * h.Num16to24)
	sumHoursWorked += 29.5 * float64(h.Num25to34)
	sumHoursWorked += float64(37 * h.Num35to39)
	sumHoursWorked += float64(40 * h.Num40)
	sumHoursWorked += 44.5 * float64(h.Num41to48)
	sumHoursWorked += float64(53 * h.Num49Plus)
	h.Average = sumHoursWorked / float64(totalNumPeople)
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
