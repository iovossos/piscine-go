package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	// If n is 1, print all single digits
	if n == 1 {
		for i := 0; i <= 9; i++ {
			z01.PrintRune(rune(i + '0')) // Print the digit
			if i != 9 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// If n is 2, print all combinations of two digits
	if n == 2 {
		for i := 0; i <= 8; i++ {
			for j := i + 1; j <= 9; j++ {
				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(j + '0'))
				if i != 8 || j != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
		return
	}

	// If n is 3, print all combinations of three digits
	if n == 3 {
		for i := 0; i <= 7; i++ {
			for j := i + 1; j <= 8; j++ {
				for k := j + 1; k <= 9; k++ {
					z01.PrintRune(rune(i + '0'))
					z01.PrintRune(rune(j + '0'))
					z01.PrintRune(rune(k + '0'))
					if i != 7 || j != 8 || k != 9 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
		z01.PrintRune('\n')
		return
	}

	// If n is 4, print all combinations of four digits
	if n == 4 {
		for i := 0; i <= 6; i++ {
			for j := i + 1; j <= 7; j++ {
				for k := j + 1; k <= 8; k++ {
					for l := k + 1; l <= 9; l++ {
						z01.PrintRune(rune(i + '0'))
						z01.PrintRune(rune(j + '0'))
						z01.PrintRune(rune(k + '0'))
						z01.PrintRune(rune(l + '0'))
						if i != 6 || j != 7 || k != 8 || l != 9 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
		return
	}

	// If n is 5, print all combinations of five digits
	if n == 5 {
		for i := 0; i <= 5; i++ {
			for j := i + 1; j <= 6; j++ {
				for k := j + 1; k <= 7; k++ {
					for l := k + 1; l <= 8; l++ {
						for m := l + 1; m <= 9; m++ {
							z01.PrintRune(rune(i + '0'))
							z01.PrintRune(rune(j + '0'))
							z01.PrintRune(rune(k + '0'))
							z01.PrintRune(rune(l + '0'))
							z01.PrintRune(rune(m + '0'))
							if i != 5 || j != 6 || k != 7 || l != 8 || m != 9 {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
		return
	}

	// If n is 6, print all combinations of six digits
	if n == 6 {
		for i := 0; i <= 4; i++ {
			for j := i + 1; j <= 5; j++ {
				for k := j + 1; k <= 6; k++ {
					for l := k + 1; l <= 7; l++ {
						for m := l + 1; m <= 8; m++ {
							for o := m + 1; o <= 9; o++ {
								z01.PrintRune(rune(i + '0'))
								z01.PrintRune(rune(j + '0'))
								z01.PrintRune(rune(k + '0'))
								z01.PrintRune(rune(l + '0'))
								z01.PrintRune(rune(m + '0'))
								z01.PrintRune(rune(o + '0'))
								if i != 4 || j != 5 || k != 6 || l != 7 || m != 8 || o != 9 {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
		return
	}

	// If n is 7, print all combinations of seven digits
	if n == 7 {
		for i := 0; i <= 3; i++ {
			for j := i + 1; j <= 4; j++ {
				for k := j + 1; k <= 5; k++ {
					for l := k + 1; l <= 6; l++ {
						for m := l + 1; m <= 7; m++ {
							for o := m + 1; o <= 8; o++ {
								for p := o + 1; p <= 9; p++ {
									z01.PrintRune(rune(i + '0'))
									z01.PrintRune(rune(j + '0'))
									z01.PrintRune(rune(k + '0'))
									z01.PrintRune(rune(l + '0'))
									z01.PrintRune(rune(m + '0'))
									z01.PrintRune(rune(o + '0'))
									z01.PrintRune(rune(p + '0'))
									if i != 3 || j != 4 || k != 5 || l != 6 || m != 7 || o != 8 || p != 9 {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
		return
	}

	// If n is 8, print all combinations of eight digits
	if n == 8 {
		for i := 0; i <= 2; i++ {
			for j := i + 1; j <= 3; j++ {
				for k := j + 1; k <= 4; k++ {
					for l := k + 1; l <= 5; l++ {
						for m := l + 1; m <= 6; m++ {
							for o := m + 1; o <= 7; o++ {
								for p := o + 1; p <= 8; p++ {
									for q := p + 1; q <= 9; q++ {
										z01.PrintRune(rune(i + '0'))
										z01.PrintRune(rune(j + '0'))
										z01.PrintRune(rune(k + '0'))
										z01.PrintRune(rune(l + '0'))
										z01.PrintRune(rune(m + '0'))
										z01.PrintRune(rune(o + '0'))
										z01.PrintRune(rune(p + '0'))
										z01.PrintRune(rune(q + '0'))
										if i != 2 || j != 3 || k != 4 || l != 5 || m != 6 || o != 7 || p != 8 || q != 9 {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
		return
	}

	// If n is 9, print all combinations of nine digits (just one: 0123456789)

	if n == 9 {
		for i := 0; i <= 1; i++ {
			for j := i + 1; j <= 2; j++ {
				for k := j + 1; k <= 3; k++ {
					for l := k + 1; l <= 4; l++ {
						for m := l + 1; m <= 5; m++ {
							for o := m + 1; o <= 6; o++ {
								for p := o + 1; p <= 7; p++ {
									for q := p + 1; q <= 8; q++ {
										for r := q + 1; r <= 9; r++ {
											z01.PrintRune(rune(i + '0'))
											z01.PrintRune(rune(j + '0'))
											z01.PrintRune(rune(k + '0'))
											z01.PrintRune(rune(l + '0'))
											z01.PrintRune(rune(m + '0'))
											z01.PrintRune(rune(o + '0'))
											z01.PrintRune(rune(p + '0'))
											z01.PrintRune(rune(q + '0'))
											z01.PrintRune(rune(r + '0'))
											if i != 1 || j != 2 || k != 3 || l != 4 || m != 5 || o != 6 || p != 7 || q != 8 || r != 9 {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
		return
	}
}