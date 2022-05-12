package card_test

import (
	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/card/french"
	"github.com/mgjules/deckr/json"
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
			c, err := card.NewCodeFromRankSuit(french.Ranks[0], french.Suits[0])
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

	Describe("Marshaling to json", func() {
		Context("with a valid code string in response", func() {
			code := "AK"

			It("should return a valid code", func() {
				type response struct {
					Code card.Code `json:"code"`
				}

				c, err := card.NewCode(code)
				Expect(err).ToNot(HaveOccurred())

				resp, err := json.Marshal(response{Code: *c})
				Expect(err).ToNot(HaveOccurred())
				Expect(string(resp)).To(Equal(`{"code":"AK"}`))
			})
		})
	})

	Describe("Unmarshaling from json", func() {
		Context("with a valid code string in request", func() {
			raw := `{"code":"AK"}`

			It("should return a valid code", func() {
				type request struct {
					Code card.Code `json:"code"`
				}

				var r request
				err := json.Unmarshal([]byte(raw), &r)
				Expect(err).ToNot(HaveOccurred())
				Expect(r.Code.String()).To(Equal("AK"))
			})
		})
	})
})
