package card_test

import (
	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/composition"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Suit", func() {
	Describe("Getting suit from code", func() {
		Context("from valid list of french suits", func() {
			comp, err := composition.FromString(composition.French)
			Expect(err).ToNot(HaveOccurred())

			suits := comp.Suits()

			Context("and a valid code", func() {
				code, err := card.CodeFromString("AS")
				Expect(err).ToNot(HaveOccurred())

				It("should return back a valid suit", func() {
					s, err := suits.SuitFromCode(*code)
					Expect(err).ToNot(HaveOccurred())
					Expect(*s).To(Equal(suits[0]))
				})
			})

			Context("and an unknown code", func() {
				code, err := card.CodeFromString("AP")
				Expect(err).ToNot(HaveOccurred())

				It("should not return any rank", func() {
					s, err := suits.SuitFromCode(*code)
					Expect(err).To(MatchError(card.ErrInvalidSuit))
					Expect(s).To(BeNil())
				})
			})
		})
	})
})
