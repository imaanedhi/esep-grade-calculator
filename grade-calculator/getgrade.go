package esepunittests

// GetGrade converts a numeric grade to a letter using inclusive thresholds.
// A ≥ 90, B ≥ 80, C ≥ 70, D ≥ 60, F < 60
func GetGrade(n int) string {
	switch {
	case n >= 90:
		return "A"
	case n >= 80:
		return "B"
	case n >= 70:
		return "C"
	case n >= 60:
		return "D"
	default:
		return "F"
	}
}
