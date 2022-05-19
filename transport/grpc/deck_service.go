package grpc

import (
	context "context"
	"errors"

	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/composition"
	"github.com/mgjules/deckr/deck"
	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo"
	"github.com/mgjules/deckr/repo/errs"
	"github.com/satori/uuid"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// DeckService is a grpc deck service.
type DeckService struct {
	log  *logger.Logger
	repo repo.Repository
	UnimplementedDeckServiceServer
}

// NewDeckService creates a new DeckService.
func NewDeckService(log *logger.Logger, repo repo.Repository) *DeckService {
	return &DeckService{
		log:  log,
		repo: repo,
	}
}

// CreateDeck creates a new full or partial deck of cards given an optional
// list of codes.
func (s *DeckService) CreateDeck(ctx context.Context, req *CreateDeckRequest) (*CreateDeckResponse, error) {
	d, err := deck.New(deck.WithComposition(*req.Comp), deck.WithCodes(req.Codes...))
	if err != nil {
		s.log.Errorf("new deck: %v", err)

		if errors.Is(err, composition.ErrUnknownComposition) ||
			errors.Is(err, card.ErrInvalidCode) ||
			errors.Is(err, card.ErrInvalidRank) ||
			errors.Is(err, card.ErrInvalidSuit) {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err := s.repo.Save(ctx, d); err != nil {
		s.log.Errorf("save deck: %v", err)

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	dc := DomainDeckToDeckClosed(d)

	return &CreateDeckResponse{
		Deck: dc,
	}, nil
}

// OpenDeck opens a deck of cards given an id.
func (s *DeckService) OpenDeck(ctx context.Context, req *OpenDeckRequest) (*OpenDeckResponse, error) {
	id := req.Id
	if _, err := uuid.FromString(id); err != nil {
		s.log.Errorf("parse id: %v", err)

		return nil, status.Errorf(codes.InvalidArgument, "invalid or missing id")
	}

	d, err := s.repo.Get(ctx, id)
	if err != nil {
		s.log.Errorf("open deck: %v", err)

		if errors.Is(err, errs.ErrDeckNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.ErrDeckNotFound.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	do := DomainDeckToDeckOpened(d)

	return &OpenDeckResponse{
		Deck: do,
	}, nil
}

// DrawCards draws cards from a deck of cards given an id and the number of
// cards.
func (s *DeckService) DrawCards(ctx context.Context, req *DrawCardsRequest) (*DrawCardsResponse, error) {
	id := req.Id
	if _, err := uuid.FromString(id); err != nil {
		s.log.Errorf("parse id: %v", err)

		return nil, status.Errorf(codes.InvalidArgument, "invalid or missing id")
	}

	num := req.Num
	if num == 0 {
		s.log.Error("parse num: invalid or missing num")

		return nil, status.Errorf(codes.InvalidArgument, "invalid or missing num")
	}

	d, err := s.repo.Get(ctx, id)
	if err != nil {
		s.log.Errorf("get deck: %v", err)

		if errors.Is(err, errs.ErrDeckNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.ErrDeckNotFound.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	drawn, err := d.Draw(int(num))
	if err != nil {
		s.log.Errorf("draw cards: %v", err)

		if errors.Is(err, deck.ErrNotEnoughCards) {
			return nil, status.Errorf(codes.InvalidArgument, deck.ErrNotEnoughCards.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err := s.repo.Save(ctx, d); err != nil {
		s.log.Errorf("save deck: %v", err)

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	cards := DomainCardsToCards(drawn)

	return &DrawCardsResponse{
		Cards: cards,
	}, nil
}

// ShuffleDeck shuffles a deck of cards given an id.
func (s *DeckService) ShuffleDeck(ctx context.Context, req *ShuffleDeckRequest) (*ShuffleDeckResponse, error) {
	id := req.Id
	if _, err := uuid.FromString(id); err != nil {
		s.log.Errorf("parse id: %v", err)

		return nil, status.Errorf(codes.InvalidArgument, "invalid or missing id")
	}

	d, err := s.repo.Get(ctx, id)
	if err != nil {
		s.log.Errorf("get deck: %v", err)

		if errors.Is(err, errs.ErrDeckNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.ErrDeckNotFound.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	d.Shuffle()

	if err := s.repo.Save(ctx, d); err != nil {
		s.log.Errorf("save deck: %v", err)

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &ShuffleDeckResponse{
		Message: "deck shuffled",
	}, nil
}

// DomainDeckToDeckClosed transforms a domain deck to a DeckClosed.
func DomainDeckToDeckClosed(d *deck.Deck) *DeckClosed {
	var dc DeckClosed
	dc.Id = d.ID()
	dc.Shuffled = d.IsShuffled()
	dc.Remaining = uint32(d.Remaining())

	return &dc
}

// DomainDeckToDeckOpened transforms a domain deck to a DeckOpened.
func DomainDeckToDeckOpened(d *deck.Deck) *DeckOpened {
	var do DeckOpened
	do.Id = d.ID()
	do.Shuffled = d.IsShuffled()
	do.Remaining = uint32(d.Remaining())
	do.Cards = DomainCardsToCards(d.Cards())

	return &do
}

// DomainCardsToCards transforms domain cards to Cards.
func DomainCardsToCards(dc []card.Card) []*Card {
	var cc []*Card
	for _, card := range dc {
		cc = append(cc, &Card{
			Value: card.Rank().String(),
			Suit:  card.Suit().String(),
			Code:  card.Code().String(),
		})
	}

	return cc
}
