package poker

type Deck interface {
	Deal() *Card
	Shuffle()
}

type Card interface {
	Suit() Suit
	Value() FaceValue
}

type Player interface {
	Role() Role // Dealer, Big Blind, Little Blind, Player
	SetRole(Role)
	AcceptCard(Card)

	Bet()
	Check()
	Call()
	Fold()
}

type Hand interface {
	Start()
	End() Player
}

type Pot interface {
}

type FaceValue int
type Suit string
type Role string
