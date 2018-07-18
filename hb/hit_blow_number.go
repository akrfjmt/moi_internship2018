package hb

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// HitBlowNumber はHB数、つまり互いに異なる数字からなるn桁の数
type HitBlowNumber struct {
	Base     uint8   // 10進数なら10
	Digits   uint8   // 4桁なら4
	Numbers  []uint8 // 実データ
	Contains []bool  // Numbersのいずれかがnであるとき、Contains[n] == true
}

func (n HitBlowNumber) Validate() error {
	if int(n.Digits) != len(n.Numbers) {
		return errors.New("digits error")
	}

	if int(n.Base) != len(n.Contains) {
		return errors.New("base error")
	}

	return nil
}

func (n HitBlowNumber) ValidateWith(m HitBlowNumber) error {
	err := n.Validate()
	if err != nil {
		return err
	}

	err = m.Validate()
	if err !=  nil {
		return err
	}

	if n.Base != m.Base {
		return errors.New("base match error")
	}

	if n.Digits != m.Digits {
		return errors.New("digits match error")
	}

	return nil
}

func (n HitBlowNumber) Product(m HitBlowNumber) (HitBlowProduct, error) {
	err := n.ValidateWith(m)
	if err != nil {
		return GetEmptyHitBlowProduct(), err
	}

	var hit uint8 = 0
	var blow uint8 = 0
	size := len(n.Numbers)

	for i := 0; i < size; i++ {
		if n.Numbers[i] == m.Numbers[i] {
			hit++
		} else if m.Contains[n.Numbers[i]] {
			blow++
		}
	}

	return HitBlowProduct{
		Base: n.Base,
		Digits: m.Digits,
		Hit: hit,
		Blow: blow,
	}, nil
}

func (n HitBlowNumber) String() string {
	s := ""
	for i := 0; i < int(n.Digits); i++ {
		s += strconv.Itoa(int(n.Numbers[i]))
	}

	return s
}

func GetEmptyHitBlowNumber() HitBlowNumber {
	return HitBlowNumber{
		Base:     0,
		Digits:   0,
		Numbers:  []uint8{},
		Contains: []bool{},
	}
}

func NewHitBlowNumber(base uint8, digits uint8, numbers []uint8) (HitBlowNumber, error) {
	if int(digits) != len(numbers) || digits < 1 || digits > 10 {
		return GetEmptyHitBlowNumber(), errors.New("digits error")
	}

	if base < 1 || base > 10 {
		return GetEmptyHitBlowNumber(), errors.New("base error")
	}

	contains := make([]bool, base)
	for i := 0; i < int(digits); i++ {
		if numbers[i] >= base {
			return GetEmptyHitBlowNumber(), errors.New("too big number: " + strconv.Itoa(int(numbers[i])))
		}

		if contains[numbers[i]] == true  {
			return GetEmptyHitBlowNumber(), errors.New("duplicated number: " + strconv.Itoa(int(numbers[i])))
		}
		contains[numbers[i]] = true
	}

	hbnum := HitBlowNumber{
		Base:     base,
		Digits:   digits,
		Numbers:  numbers,
		Contains: contains,
	}

	return hbnum, nil
}

func CreateAllHitBlowNumbers(base uint8, digits uint8) []HitBlowNumber {
	if base < digits {
		fmt.Errorf("base digits error")
		os.Exit(1)
	}

	switch digits {
	case 3:
		return createAllHitBlowNumbersDigits3(base)

	case 10:
		return createAllHitBlowNumbersDigits10(base)
	default:
		fmt.Errorf("unsupported digits")
		os.Exit(1)
	}

	return []HitBlowNumber{}
}

func createAllHitBlowNumbersDigits3(base uint8) []HitBlowNumber {
	if base < 3 {
		fmt.Errorf("base digits error")
		os.Exit(1)
	}

	candidates := []uint8{}

	for i := uint8(0); i < base; i++ {
		candidates = append(candidates, i)
	}

	questions := []HitBlowNumber{}

	for k := uint8(0); k < base; k++ {

		for j := uint8(0); j < base; j++ {
			if k == j {
				continue
			}

			for i := uint8(0); i < base; i++ {
				if k == i || j == i {
					continue
				}

				question, _ := NewHitBlowNumber(base, 3, []uint8{k,j,i})
				questions = append(questions, question)
			}
		}
	}

	return questions
}

func createAllHitBlowNumbersDigits10(base uint8) []HitBlowNumber {
	if base > 10 {
		fmt.Errorf("base digits error")
		os.Exit(1)
	}

	candidates := []uint8{}

	for i := uint8(0); i < base; i++ {
		candidates = append(candidates, i)
	}

	hbNumers := []HitBlowNumber{}

	for d0 := uint8(0); d0 < base; d0++ {
		for d1 := uint8(0); d1 < base; d1++ {
			if d1 == d0 {
				continue
			}
			for d2 := uint8(0); d2 < base; d2++ {
				if d2 == d0 || d2 == d1 {
					continue
				}
				for d3 := uint8(0); d3 < base; d3++ {
					if d3 == d0 || d3 == d1 || d3 == d2 {
						continue
					}
					for d4 := uint8(0); d4 < base; d4++ {
						if d4 == d0 || d4 == d1 || d4 == d2 ||d4 == d3 {
							continue
						}
						for d5 := uint8(0); d5 < base; d5++ {
							if d5 == d0 || d5 == d1 || d5 == d2 ||d5 == d3 ||d5 == d4 {
								continue
							}
							for d6 := uint8(0); d6 < base; d6++ {
								if d6 == d0 || d6 == d1 || d6 == d2 ||d6 == d3 ||d6 == d4 || d6 == d5 {
									continue
								}
								for d7 := uint8(0); d7 < base; d7++ {
									if d7 == d0 || d7 == d1 || d7 == d2 ||d7 == d3 ||d7 == d4 || d7 == d5 || d7 == d6 {
										continue
									}
									for d8 := uint8(0); d8 < base; d8++ {
										if d8 == d0 || d8 == d1 || d8 == d2 ||d8 == d3 ||d8 == d4 || d8 == d5 || d8 == d6 || d8 == d7 {
											continue
										}
										for d9 := uint8(0); d9 < base; d9++ {
											if d9 == d0 || d9 == d1 || d9 == d2 || d9 == d3 || d9 == d4 || d9 == d5 || d9 == d6 || d9 == d7 || d9 == d8 {
												continue
											}
											hbNumer, _ := NewHitBlowNumber(base, 10, []uint8{d0, d1, d2, d3, d4, d5, d6, d7, d8, d9})
											hbNumers = append(hbNumers, hbNumer)
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

	return hbNumers
}
