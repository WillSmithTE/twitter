package vaccineDemographics

type Database struct {
	Data []*AreaData
}

type AreaData struct {
	Area                 Area
	VaccinatedPercentage float64
	IncidenceRate        float64
}
