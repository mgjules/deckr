package card_test

import (
	"io/ioutil"

	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/card/french"
	"github.com/mgjules/deckr/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Card", func() {
	Describe("Creating a list of cards", func() {
		Context("from valid list of french ranks and suits", func() {
			type expected struct {
				Cards []struct {
					Rank string `json:"rank"`
					Suit string `json:"suit"`
					Code string `json:"code"`
				} `json:"cards"`
			}

			ranks := french.Ranks
			suits := french.Suits

			Context("and with no codes", func() {
				It("should return back a full list of cards in the correct order", func() {
					cards, err := card.NewCards(ranks, suits)
					Expect(err).ToNot(HaveOccurred())
					Expect(cards).To(HaveLen(52))

					raw, err := ioutil.ReadFile(testdataDir + "/full_french_cards.json")
					Expect(err).ToNot(HaveOccurred())

					var exp expected
					err = json.Unmarshal(raw, &exp)
					Expect(err).ToNot(HaveOccurred())

					Expect(cards).To(HaveLen(len(exp.Cards)))

					for i, c := range cards {
						Expect(c.Rank().String()).To(Equal(exp.Cards[i].Rank))
						Expect(c.Suit().String()).To(Equal(exp.Cards[i].Suit))
						Expect(c.Code().String()).To(Equal(exp.Cards[i].Code))
					}
				})
			})

			Context("and with specific codes", func() {
				codes, err := card.NewCodes("AS", "KD", "AC", "2C", "KH")
				Expect(err).ToNot(HaveOccurred())

				It("should return back a partial list of cards in same order as codes", func() {
					cards, err := card.NewCards(ranks, suits, codes...)
					Expect(err).ToNot(HaveOccurred())
					Expect(cards).To(HaveLen(5))

					raw, err := ioutil.ReadFile(testdataDir + "/partial_french_cards.json")
					Expect(err).ToNot(HaveOccurred())

					var exp expected
					err = json.Unmarshal(raw, &exp)
					Expect(err).ToNot(HaveOccurred())

					Expect(cards).To(HaveLen(len(exp.Cards)))

					for i, c := range cards {
						Expect(c.Rank().String()).To(Equal(exp.Cards[i].Rank))
						Expect(c.Suit().String()).To(Equal(exp.Cards[i].Suit))
						Expect(c.Code().String()).To(Equal(exp.Cards[i].Code))
					}
				})
			})
		})
	})
})
