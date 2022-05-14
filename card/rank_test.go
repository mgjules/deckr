package card_test

import (
	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/card/french"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rank", func() {
	Describe("Getting rank from code", func() {
		Context("from valid list of french ranks", func() {
			ranks := french.Composition.Ranks()

			Context("and a valid code", func() {
				code, err := card.NewCode("AS")
				Expect(err).ToNot(HaveOccurred())

				It("should return back a valid rank", func() {
					r, ok := ranks.RankFromCode(*code)
					Expect(ok).To(BeTrue())
					Expect(*r).To(Equal(ranks[0]))
				})
			})

			Context("and an unknown code", func() {
				code, err := card.NewCode("PS")
				Expect(err).ToNot(HaveOccurred())

				It("should not return any rank", func() {
					r, ok := ranks.RankFromCode(*code)
					Expect(ok).To(BeFalse())
					Expect(r).To(BeNil())
				})
			})
		})
	})
})
