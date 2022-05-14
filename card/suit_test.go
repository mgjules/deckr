package card_test

import (
	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/card/french"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Suit", func() {
	Describe("Getting suit from code", func() {
		Context("from valid list of french suits", func() {
			suits := french.Composition.Suits()

			Context("and a valid code", func() {
				code, err := card.NewCode("AS")
				Expect(err).ToNot(HaveOccurred())

				It("should return back a valid suit", func() {
					s, ok := suits.SuitFromCode(*code)
					Expect(ok).To(BeTrue())
					Expect(*s).To(Equal(suits[0]))
				})
			})

			Context("and an unknown code", func() {
				code, err := card.NewCode("AP")
				Expect(err).ToNot(HaveOccurred())

				It("should not return any rank", func() {
					s, ok := suits.SuitFromCode(*code)
					Expect(ok).To(BeFalse())
					Expect(s).To(BeNil())
				})
			})
		})
	})
})
