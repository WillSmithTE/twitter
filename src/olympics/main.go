package olympics

func Main() {
	countries := []*Country{}
	
}

func getAndServeSongs() {
}

type Country struct {
	Name        string
	YearResults []YearResult
}

type YearResult struct {
	year      int
	NumBronze int
	NumSilver int
	NumGold   int
}

func (y *YearResult) getTotal() int {
	return y.NumBronze + y.NumSilver + y.NumGold
}
