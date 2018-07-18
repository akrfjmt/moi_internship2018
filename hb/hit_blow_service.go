package hb

import (
	"github.com/akrfjmt/moi_internship2018/web"
	"errors"
)

type HitBlowService struct {
	Base uint8
	Digits uint8
}

func (s HitBlowService) CreateQuestionRequest(hbNumber HitBlowNumber) web.QuestionRequest {
	question := web.QuestionRequest{
		Question: hbNumber.String(),
	}

	return question
}

func (s HitBlowService) CreateHitBlowProduct(questionResponse web.QuestionResponse) HitBlowProduct {
	hitBlowProduct := HitBlowProduct{
		Base: s.Base,
		Digits: s.Digits,
		Hit: uint8(questionResponse.Hit),
		Blow: uint8(questionResponse.Blow),
	}

	return hitBlowProduct
}

func (s HitBlowService) FilterQuestions(lastQuestion HitBlowNumber, lastProduct HitBlowProduct, questions []HitBlowNumber) []HitBlowNumber {
	filteredHbNumbers := []HitBlowNumber{}

	for i := 0; i < len(questions); i++ {
		question := questions[i]
		product, _ := lastQuestion.Product(question)
		if lastProduct.Equals(product) {
			filteredHbNumbers = append(filteredHbNumbers, question)
		}
	}

	return filteredHbNumbers
}

// HitBlowProductが何種類の値を取り得るかを返す
func (s HitBlowService) ProductSpace() uint8 {
	hitBlowProduct := HitBlowProduct{
		Base: s.Base,
		Digits: s.Digits,
		Hit: s.Digits,
		Blow: 0,
	}
	value, _ := hitBlowProduct.Value()
	return value + 1
}

func (s HitBlowService) AllProducts() map[uint8]HitBlowProduct {
	products := map[uint8]HitBlowProduct{}

	for hit := uint8(0); hit <= s.Digits; hit++ {
		// blowの最大値。例えば、Digits = 3でhit = 1のとき、blowの最大値は2となる。
		maxBlow := s.Digits - hit

		for blow := uint8(0); blow <= maxBlow; blow++ {
			if hit == s.Digits - 1 && blow == 1 {
				continue
			}

			product := HitBlowProduct{
				Base: s.Base,
				Digits: s.Digits,
				Hit: hit,
				Blow: blow,
			}

			productValue, _ := product.Value()
			products[productValue] = product
		}
	}

	return products
}

func (s HitBlowService) ResolveHitBlowProduct(value uint8) (*HitBlowProduct, error) {
	allProducts := s.AllProducts()
	if _, ok := allProducts[value]; ok {

		product := allProducts[value]
		return &product, nil
	}

	return nil, errors.New("fatal error.")
}

func (s HitBlowService) CreateAllHitBlowNumbers() []HitBlowNumber {
	return CreateAllHitBlowNumbers(s.Base, s.Digits)
}
