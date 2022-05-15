package deck_test

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/composition"
	"github.com/mgjules/deckr/deck"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const testdataDir = "testdata"

var _ = Describe("Deck", func() {
	Describe("Creating a deck", func() {
		Context("from valid list of french cards", func() {
			type expected struct {
				Cards []struct {
					Rank string `json:"rank"`
					Suit string `json:"suit"`
					Code string `json:"code"`
				} `json:"cards"`
			}

			Context("and with no codes", func() {
				It("should return back a full list of cards in the correct order", func() {
					d, err := deck.New(deck.WithComposition(composition.French))
					Expect(err).ToNot(HaveOccurred())
					Expect(d.Remaining()).To(Equal(52))

					raw, err := ioutil.ReadFile(testdataDir + "/full_french_cards.json")
					Expect(err).ToNot(HaveOccurred())

					var exp expected
					err = json.Unmarshal(raw, &exp)
					Expect(err).ToNot(HaveOccurred())

					Expect(d.Remaining()).To(Equal(len(exp.Cards)))

					for i, c := range d.Cards() {
						Expect(c.Rank().String()).To(Equal(exp.Cards[i].Rank))
						Expect(c.Suit().String()).To(Equal(exp.Cards[i].Suit))
						Expect(c.Code().String()).To(Equal(exp.Cards[i].Code))
					}
				})
			})

			Context("and with specific codes", func() {
				It("should return back a partial list of cards in same order as codes", func() {
					d, err := deck.New(
						deck.WithComposition(composition.French),
						deck.WithCodes("AS", "KD", "AC", "2C", "KH"),
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(d.Remaining()).To(Equal(5))

					raw, err := ioutil.ReadFile(testdataDir + "/partial_french_cards.json")
					Expect(err).ToNot(HaveOccurred())

					var exp expected
					err = json.Unmarshal(raw, &exp)
					Expect(err).ToNot(HaveOccurred())

					Expect(d.Remaining()).To(Equal(len(exp.Cards)))

					for i, c := range d.Cards() {
						Expect(c.Rank().String()).To(Equal(exp.Cards[i].Rank))
						Expect(c.Suit().String()).To(Equal(exp.Cards[i].Suit))
						Expect(c.Code().String()).To(Equal(exp.Cards[i].Code))
					}
				})
			})
		})
	})

	Describe("Shuffling a deck of french cards", func() {
		Context("which is full", func() {
			d, err := deck.New(deck.WithComposition(composition.French))
			Expect(err).ToNot(HaveOccurred())
			Expect(d.IsShuffled()).To(BeFalse())

			var original []card.Card
			original = append(original, d.Cards()...)

			It("should return a shuffled deck", func() {
				d.Shuffle()

				Expect(d.Remaining()).To(Equal(len(original)))
				Expect(d.Cards()).ToNot(Equal(original))
				Expect(d.IsShuffled()).To(BeTrue())
			})
		})

		Context("which is partial", func() {
			d, err := deck.New(
				deck.WithComposition(composition.French),
				deck.WithCodes("AS", "KD", "AC", "2C", "KH"),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(d.IsShuffled()).To(BeFalse())

			var original []card.Card
			original = append(original, d.Cards()...)

			It("should return a shuffled deck", func() {
				d.Shuffle()

				Expect(d.Cards()).To(HaveLen(len(original)))
				Expect(d.Cards()).ToNot(Equal(original))
				Expect(d.IsShuffled()).To(BeTrue())
			})
		})
	})

	Describe("Drawing cards from a deck of french cards", func() {
		Context("which is full", func() {
			d, err := deck.New(deck.WithComposition(composition.French))
			Expect(err).ToNot(HaveOccurred())
			Expect(d.Remaining()).To(Equal(52))

			Context("with enough cards", func() {
				It("should return the top cards in the stack", func() {
					drawn, err := d.Draw(5)
					Expect(err).ToNot(HaveOccurred())
					Expect(drawn).To(HaveLen(5))
					Expect(d.Remaining()).To(Equal(47))
				})
			})

			Context("with enough cards again", func() {
				It("should return the top cards in the stack", func() {
					drawn, err := d.Draw(10)
					Expect(err).ToNot(HaveOccurred())
					Expect(drawn).To(HaveLen(10))
					Expect(d.Remaining()).To(Equal(37))
				})
			})
		})

		Context("which is partial", func() {
			d, err := deck.New(
				deck.WithComposition(composition.French),
				deck.WithCodes("AS", "KD", "AC", "2C", "KH"),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(d.Remaining()).To(Equal(5))

			Context("with enough cards", func() {
				It("should return the top cards in the stack", func() {
					drawn, err := d.Draw(3)
					Expect(err).ToNot(HaveOccurred())
					Expect(drawn).To(HaveLen(3))
					Expect(d.Remaining()).To(Equal(2))
				})
			})

			Context("without enough cards", func() {
				It("should return an error", func() {
					drawn, err := d.Draw(3)
					Expect(err).To(MatchError(deck.ErrNotEnoughCards))
					Expect(drawn).To(HaveLen(0))
					Expect(d.Remaining()).To(Equal(2))
				})
			})
		})
	})
})
