package card_test

import (
	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/composition"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rank", func() {
	Describe("Getting rank from code", func() {
		Context("from valid list of french ranks", func() {
			comp, err := composition.FromString(composition.French)
			Expect(err).ToNot(HaveOccurred())

			ranks := comp.Ranks()

			Context("and a valid code", func() {
				code, err := card.CodeFromString("AS")
				Expect(err).ToNot(HaveOccurred())

				It("should return back a valid rank", func() {
					r, err := ranks.RankFromCode(*code)
					Expect(err).ToNot(HaveOccurred())
					Expect(*r).To(Equal(ranks[0]))
				})
			})

			Context("and an unknown code", func() {
				code, err := card.CodeFromString("PS")
				Expect(err).ToNot(HaveOccurred())

				It("should not return any rank", func() {
					r, err := ranks.RankFromCode(*code)
					Expect(err).To(MatchError(card.ErrInvalidRank))
					Expect(r).To(BeNil())
				})
			})
		})
	})
})
