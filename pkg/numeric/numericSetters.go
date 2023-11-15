package numeric

// SetNumericUpdateHandler returns a function that sets the update handler for a Numeric widget.
// The handler function is expected to return an integer and an error, which will be used to update the Numeric widget's value.
func SetNumericUpdateHandler(handler func() (int, error)) func(
	f Numeric,
) {
	return func(f Numeric) {
		f.setUpdateHandler(handler)
	}
}

// SetNumericInitValue returns a function that sets the initial value for a Numeric widget.
// The value is provided as an integer and is intended to be the starting point for the widget's displayed value.
func SetNumericInitValue(value int) func(
	f Numeric,
) {
	return func(f Numeric) {
		f.setInitialValue(value)
	}
}

// SetNumericUnit returns a function that sets the measurement unit for a Numeric widget.
// The unit is a string that represents the unit of measurement to be displayed alongside the Numeric value.
func SetNumericUnit(unit string) func(
	f Numeric,
) {
	return func(f Numeric) {
		f.setUnit(unit)
	}
}

// SetNumericUnitAfter returns a function that configures a Numeric widget to display its unit after the numeric value.
// This function does not take any arguments and triggers the 'withUnitAfter' method on the Numeric widget.
func SetNumericUnitAfter() func(
	f Numeric,
) {
	return func(f Numeric) {
		f.withUnitAfter()
	}
}
