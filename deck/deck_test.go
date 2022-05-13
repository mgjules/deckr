package deck_test

import (
	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/card/french"
	"github.com/mgjules/deckr/deck"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deck", func() {
	Describe("Shuffling a deck of french cards", func() {
		ranks := french.Ranks
		suits := french.Suits

		Context("which is full", func() {
			cards, err := card.NewCards(ranks, suits)
			Expect(err).ToNot(HaveOccurred())

			var original []card.Card
			original = append(original, cards...)

			It("should return a shuffled deck", func() {
				deck, err := deck.New(deck.WithCards(cards...))
				Expect(err).ToNot(HaveOccurred())
				Expect(deck.IsShuffled()).To(BeFalse())

				deck.Shuffle()

				Expect(deck.Remaining()).To(Equal(len(original)))
				Expect(deck.Cards()).ToNot(Equal(original))
				Expect(deck.IsShuffled()).To(BeTrue())
			})
		})

		Context("which is partial", func() {
			codes, err := card.NewCodes("AS", "KD", "AC", "2C", "KH")
			Expect(err).ToNot(HaveOccurred())

			cards, err := card.NewCards(ranks, suits, codes...)
			Expect(err).ToNot(HaveOccurred())

			var original []card.Card
			original = append(original, cards...)

			It("should return a shuffled deck", func() {
				deck, err := deck.New(deck.WithCards(cards...))
				Expect(err).ToNot(HaveOccurred())
				Expect(deck.IsShuffled()).To(BeFalse())

				deck.Shuffle()

				Expect(deck.Cards()).To(HaveLen(len(original)))
				Expect(deck.Cards()).ToNot(Equal(original))
				Expect(deck.IsShuffled()).To(BeTrue())
			})
		})
	})

	Describe("Drawing cards from a deck of french cards", func() {
		ranks := french.Ranks
		suits := french.Suits

		Context("which is full", func() {
			cards, err := card.NewCards(ranks, suits)
			Expect(err).ToNot(HaveOccurred())

			deck, err := deck.New(deck.WithCards(cards...))
			Expect(err).ToNot(HaveOccurred())
			Expect(cards).To(HaveLen(52))

			Context("with enough cards", func() {
				It("should return the top cards in the stack", func() {
					drawn, err := deck.Draw(5)
					Expect(err).ToNot(HaveOccurred())
					Expect(drawn).To(HaveLen(5))
					Expect(deck.Remaining()).To(Equal(47))
				})
			})

			Context("with enough cards again", func() {
				It("should return the top cards in the stack", func() {
					drawn, err := deck.Draw(10)
					Expect(err).ToNot(HaveOccurred())
					Expect(drawn).To(HaveLen(10))
					Expect(deck.Remaining()).To(Equal(37))
				})
			})
		})

		Context("which is partial", func() {
			codes, err := card.NewCodes("AS", "KD", "AC", "2C", "KH")
			Expect(err).ToNot(HaveOccurred())

			cards, err := card.NewCards(ranks, suits, codes...)
			Expect(err).ToNot(HaveOccurred())

			deck, err := deck.New(deck.WithCards(cards...))
			Expect(err).ToNot(HaveOccurred())
			Expect(cards).To(HaveLen(5))

			Context("with enough cards", func() {
				It("should return the top cards in the stack", func() {
					drawn, err := deck.Draw(3)
					Expect(err).ToNot(HaveOccurred())
					Expect(drawn).To(HaveLen(3))
					Expect(deck.Remaining()).To(Equal(2))
				})
			})

			Context("without enough cards", func() {
				It("should return an error", func() {
					_, err := deck.Draw(3)
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
