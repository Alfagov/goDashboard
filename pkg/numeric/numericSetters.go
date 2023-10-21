package numeric

func SetNumericUpdateHandler(handler func() (int, error)) func(
	f Numeric,
) {
	return func(f Numeric) {
		f.setUpdateHandler(handler)
	}
}

func SetNumericInitValue(value int) func(
	f Numeric,
) {
	return func(f Numeric) {
		f.setInitialValue(value)
	}
}

func SetNumericUnit(unit string) func(
	f Numeric,
) {
	return func(f Numeric) {
		f.setUnit(unit)
	}
}

func SetNumericUnitAfter() func(
	f Numeric,
) {
	return func(f Numeric) {
		f.withUnitAfter()
	}
}
