package models

type Basket struct {
	Diets map[string]Diet `json:"diets" bson:"diets"`
}

func NewBasket() *Basket {
	basket := &Basket{}
	basket.Diets = make(map[string]Diet)
	return basket
}

func (b *Basket) Delete(id string) {
	delete(b.Diets, id)
}
