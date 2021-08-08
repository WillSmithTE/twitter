package vaccineDemographics

import (
	"log"
	"math"

	"github.com/willsmithte/twitter/src/util"
)

type ComputedData struct {
	ReligionCorrelations map[string]float64
}

func NewComputedData(data []*AreaData) *ComputedData {
	return &ComputedData{ReligionCorrelations: calculateReligionCorrelations(data)}
}

func calculateReligionCorrelations(data []*AreaData) map[string]float64 {
	correlations := make(map[string]float64)
	for i := 0; i < len(data); i++ {
		for religionName := range data[i].Area.GetReligionPercentages() {
			_, ok := correlations[religionName]
			if !ok {
				var religionPct, pctVaxxed []float64
				for j := i; j < len(data); j++ {
					pct, areaHasReligion := data[j].Area.GetReligionPercentages()[religionName]
					if areaHasReligion {
						religionPct = append(religionPct, pct)
						pctVaxxed = append(pctVaxxed, data[j].CovidVaccine.Num1Dose)
					}
				}
				corrCoeff := correlationCoefficient(religionPct, pctVaxxed)
				if math.IsNaN(corrCoeff) {
					log.Printf("Correlation coefficient is NaN for %v %v\n", religionName, data[i].Area.Name4)
				} else {
					correlations[religionName] = corrCoeff
				}
			}
		}
	}
	util.PrintJson(correlations)
	return correlations
}

// from https://socketloop.com/tutorials/golang-find-correlation-coefficient-example
func correlationCoefficient(X []float64, Y []float64) float64 {
	n := float64(len(X))

	sum_X := 0.0
	sum_Y := 0.0
	sum_XY := 0.0
	squareSum_X := 0.0
	squareSum_Y := 0.0

	for i := 0; i < int(n); i++ {
		sum_X = sum_X + X[i]
		sum_Y = sum_Y + Y[i]
		sum_XY = sum_XY + X[i]*Y[i]
		squareSum_X = squareSum_X + X[i]*X[i]
		squareSum_Y = squareSum_Y + Y[i]*Y[i]
	}

	corr := float64((n*sum_XY - sum_X*sum_Y)) /
		(math.Sqrt(float64((n*squareSum_X - sum_X*sum_X) * (n*squareSum_Y - sum_Y*sum_Y))))

	return corr
}
