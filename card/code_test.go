package card_test

import (
	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/composition"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Code", func() {
	Describe("Creating a new code", func() {
		Context("with a valid code string", func() {
			c, err := card.NewCode("AS")
			Expect(err).ToNot(HaveOccurred())

			It("should print back a valid code", func() {
				Expect(c.String()).To(Equal("AS"))
			})

			It("should print back a valid rank and suit", func() {
				Expect(c.Rank()).To(Equal("A"))
				Expect(c.Suit()).To(Equal("S"))
			})
		})

		Context("with a valid french rank and suit", func() {
			comp, err := composition.ParseFromString(composition.French)
			Expect(err).ToNot(HaveOccurred())

			c, err := card.NewCodeFromRankSuit(comp.Ranks()[0], comp.Suits()[0])
			Expect(err).ToNot(HaveOccurred())

			It("should print back a valid code", func() {
				Expect(c.String()).To(Equal("AS"))
			})

			It("should print back a valid rank and suit", func() {
				Expect(c.Rank()).To(Equal("A"))
				Expect(c.Suit()).To(Equal("S"))
			})
		})
	})

	Describe("Creating a list of codes", func() {
		Context("with a valid list of codes string", func() {
			ak, err := card.NewCode("AK")
			Expect(err).ToNot(HaveOccurred())

			twos, err := card.NewCode("2S")
			Expect(err).ToNot(HaveOccurred())

			cc, err := card.NewCodes("AK", "2S")
			Expect(err).ToNot(HaveOccurred())

			It("should return a valid list of codes", func() {
				Expect(cc).To(Equal([]card.Code{*ak, *twos}))
			})
		})
	})
})
