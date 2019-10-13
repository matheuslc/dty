package receipts

// Receipt defines the main product
// A receipt have an amount of ingredients and directions
type Receipt struct {
	ingredients []Ingredient
	name        string
	portions    int
	direction   Direction
}
