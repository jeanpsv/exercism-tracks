package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var decision string
	switch {
	case card1 == "ace" && card2 == "ace":
		decision = "P"
	case (ParseCard(card1)+ParseCard(card2) == 21) && ParseCard(dealerCard) < 10:
		decision = "W"
	case (ParseCard(card1)+ParseCard(card2) == 21) && ParseCard(dealerCard) >= 10:
		decision = "S"
	case ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20:
		decision = "S"
	case (ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16) && ParseCard(dealerCard) < 7:
		decision = "S"
	case (ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16) && ParseCard(dealerCard) >= 7:
		decision = "H"
	case ParseCard(card1)+ParseCard(card2) <= 11:
		decision = "H"
	}
	return decision
}
