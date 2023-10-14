package widgets

func SetNumericUpdateHandler(handler func() (int, error)) func(
	f NumericWidget,
) {
	return func(f NumericWidget) {
		f.setUpdateHandler(handler)
	}
}

func SetNumericInitValue(value int) func(
	f NumericWidget,
) {
	return func(f NumericWidget) {
		f.setInitialValue(value)
	}
}

func SetNumericUnit(unit string) func(
	f NumericWidget,
) {
	return func(f NumericWidget) {
		f.setUnit(unit)
	}
}

func SetNumericUnitAfter() func(
	f NumericWidget,
) {
	return func(f NumericWidget) {
		f.withUnitAfter()
	}
}
