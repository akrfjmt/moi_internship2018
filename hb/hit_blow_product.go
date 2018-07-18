package hb

import (
	"errors"
	"strconv"
)

type HitBlowProduct struct {
	Base uint8
	Digits uint8
	Hit uint8
	Blow uint8
}

func (p HitBlowProduct) Value() (uint8, error) {
	if p.Base < 0 {
		return 0, errors.New("base is smaller than 0")
	}

	if p.Hit < 0 {
		return 0, errors.New("h is smaller than 0")
	}

	if p.Hit > p.Digits {
		return 0, errors.New("h is bigger than digits")
	}

	if p.Blow < 0 {
		return 0, errors.New("b is smaller than 0")
	}

	if p.Blow > p.Digits {
		return 0, errors.New("h is bigger than digits")
	}

	num := p.Hit - 1

	if p.Hit == p.Digits - 1 && p.Blow == 1 {
		// 例えば、Digitsが3だった場合に2hit 1blowはありえない。
		return 0, errors.New("bad request")
	}

	if p.Hit == p.Digits {
		return p.Digits * p.Hit - num * (num + 1) / 2 + p.Hit + p.Blow - 1, nil
	}

	return p.Digits * p.Hit - num * (num + 1) / 2 + p.Hit + p.Blow, nil
}

func (p HitBlowProduct) String() string {
	value, _ := p.Value()
	return strconv.Itoa(int(value)) + ":" + strconv.Itoa(int(p.Hit)) + "H" + strconv.Itoa(int(p.Blow)) + "B"
}

func (p HitBlowProduct) Equals(q HitBlowProduct) bool {
	return p.Hit == q.Hit && p.Blow == q.Blow
}

func GetEmptyHitBlowProduct() HitBlowProduct {
	return HitBlowProduct{
		Hit: 0,
		Blow: 0,
	}
}
