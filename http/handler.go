package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/card/french"
	"github.com/mgjules/deckr/deck"
	"github.com/mgjules/deckr/docs"
	"github.com/mgjules/deckr/repo/inmemory"
	"github.com/satori/uuid"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// handleHealthCheck godoc
// @Summary      Health Check
// @Description  checks if server is running
// @Tags         core
// @Produce      json
// @Success      200  {string}  I'm  alive!
// @Router       / [get]
func (Server) handleHealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, Success{Message: "I'm alive!"})
	}
}

// handleVersion godoc
// @Summary      Version
// @Description  checks the server's version
// @Tags         core
// @Produce      json
// @Success      200  {object}  build.Info
// @Router       /version [get]
func (s *Server) handleVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, s.build)
	}
}

func (Server) handleSwagger() gin.HandlerFunc {
	docs.SwaggerInfo.BasePath = "/"

	url := ginSwagger.URL("/swagger/doc.json")

	return ginSwagger.WrapHandler(swaggerFiles.Handler, url)
}

// handleCreateDeck godoc
// @Summary      creates a new deck of cards
// @Description  creates a new full or partial deck of cards given an optional list of codes
// @Tags         deck
// @Produce      json
// @Param        codes  query     []string  false  "list of codes"  example(AS, 2C, 3D, 4H, 5S)
// @Success      201    {object}  http.DeckClosed
// @Failure      400    {object}  http.Error
// @Failure      500    {object}  http.Error
// @Router       /decks [post]
func (s *Server) handleCreateDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := c.QueryArray("codes")

		codes, err := card.NewCodes(cc...)
		if err != nil {
			s.log.Errorf("new codes: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, Error{err.Error()})

			return
		}

		cards, err := card.NewCards(french.Composition, codes...)
		if err != nil {
			s.log.Errorf("new cards: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		d, err := deck.New(deck.WithCards(cards...))
		if err != nil {
			s.log.Errorf("new deck: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		rd := DomainDeckToRepoDeck(d)
		if err := s.repo.Save(c, rd); err != nil {
			s.log.Errorf("save deck: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		dc := DomainDeckToDeckClosed(d)

		c.JSON(http.StatusCreated, dc)
	}
}

// handleOpenDeck godoc
// @Summary      opens a deck of cards
// @Description  opens a deck of cards given an id
// @Tags         deck
// @Produce      json
// @Param        id   path      string  true  "id of deck"       example(9302b603-13bb-5275-a3b9-5fcefafa34e0)
// @Success      200  {object}  http.DeckOpened
// @Failure      400  {object}  http.Error
// @Failure      404  {object}  http.Error
// @Failure      500  {object}  http.Error
// @Router       /decks/{id} [get]
func (s *Server) handleOpenDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if _, err := uuid.FromString(id); err != nil {
			s.log.Errorf("parse id: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, Error{"invalid or missing id"})

			return
		}

		rd, err := s.repo.Get(c, id)
		if err != nil {
			s.log.Errorf("open deck: %v", err)

			if errors.Is(err, inmemory.ErrDeckNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, Error{inmemory.ErrDeckNotFound.Error()})

				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		do := RepoDeckToDeckOpened(rd)

		c.JSON(http.StatusOK, do)
	}
}

// handleDrawCards godoc
// @Summary      draws cards from a deck of cards
// @Description  draws cards from a deck of cards given an id and the number of cards
// @Tags         deck
// @Produce      json
// @Param        id   path      string  true  "id of deck"  example(9302b603-13bb-5275-a3b9-5fcefafa34e0)
// @Param        num  query     int     true  "number of cards"  example(5)
// @Success      200  {object}  http.Cards
// @Failure      400  {object}  http.Error
// @Failure      404  {object}  http.Error
// @Failure      500  {object}  http.Error
// @Router       /decks/{id}/draw [get]
func (s *Server) handleDrawCards() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if _, err := uuid.FromString(id); err != nil {
			s.log.Errorf("parse id: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, Error{"invalid or missing id"})

			return
		}

		num, err := strconv.Atoi(c.Query("num"))
		if err != nil || num == 0 {
			s.log.Errorf("parse num: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, Error{"invalid or missing num"})

			return
		}

		rd, err := s.repo.Get(c, id)
		if err != nil {
			s.log.Errorf("get deck: %v", err)

			if errors.Is(err, inmemory.ErrDeckNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, Error{inmemory.ErrDeckNotFound.Error()})

				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		d, err := RepoDeckToDomainDeck(rd)
		if err != nil {
			s.log.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		drawn, err := d.Draw(num)
		if err != nil {
			s.log.Errorf("draw cards: %v", err)

			if errors.Is(err, deck.ErrNotEnoughCards) {
				c.AbortWithStatusJSON(http.StatusBadRequest, Error{deck.ErrNotEnoughCards.Error()})

				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		rd = DomainDeckToRepoDeck(d)
		if err := s.repo.Save(c, rd); err != nil {
			s.log.Errorf("save deck: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		cards := DomainCardsToCards(drawn)

		c.JSON(http.StatusOK, cards)
	}
}

// handleShuffleDeck godoc
// @Summary      shuffle a deck of cards
// @Description  shuffle a deck of cards given an id
// @Tags         deck
// @Produce      json
// @Param        id   path      string  true  "id of deck"  example(9302b603-13bb-5275-a3b9-5fcefafa34e0)
// @Success      200  {string}  deck    shuffled
// @Failure      400  {object}  http.Error
// @Failure      404  {object}  http.Error
// @Failure      500  {object}  http.Error
// @Router       /decks/{id}/shuffle [post]
func (s *Server) handleShuffleDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if _, err := uuid.FromString(id); err != nil {
			s.log.Errorf("parse id: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, Error{"invalid or missing id"})

			return
		}

		rd, err := s.repo.Get(c, id)
		if err != nil {
			s.log.Errorf("get deck: %v", err)

			if errors.Is(err, inmemory.ErrDeckNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, Error{inmemory.ErrDeckNotFound.Error()})

				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		d, err := RepoDeckToDomainDeck(rd)
		if err != nil {
			s.log.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		d.Shuffle()

		rd = DomainDeckToRepoDeck(d)
		if err := s.repo.Save(c, rd); err != nil {
			s.log.Errorf("save deck: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, Error{err.Error()})

			return
		}

		c.JSON(http.StatusOK, Success{"deck shuffled"})
	}
}
