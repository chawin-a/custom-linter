package complexity

func highComplexity() { // want "calculated cyclomatic complexity for function"
	i := 1
	if i > 2 {
		if i > 2 && i > 3 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
	} else {
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
	}

	if i > 2 {
	}
}

func noComplexity() {} // want "calculated cyclomatic complexity for function"
