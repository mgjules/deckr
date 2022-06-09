package postgres

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	mpgx "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/mgjules/deckr/deck"
	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo/errs"
)

//go:embed migrations/*.sql
var migrations embed.FS

const (
	table   = "decks"
	version = 1
)

// Repository is a PostgreSQL implementation of the deckr.Repository interface.
type Repository struct {
	log *logger.Logger
	pgx *pgx.Conn
	db  *sql.DB
	sb  sq.StatementBuilderType
}

// NewRepository creates a new PostgreSQL repository.
func NewRepository(ctx context.Context, uri string, log *logger.Logger) (*Repository, error) {
	db, err := sql.Open("pgx", uri)
	if err != nil {
		return nil, fmt.Errorf("open postgres db: %w", err)
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres db: %w", err)
	}

	conn, err := stdlib.AcquireConn(db)
	if err != nil {
		return nil, fmt.Errorf("acquire pgx connection: %w", err)
	}

	return &Repository{
		log: log,
		pgx: conn,
		db:  db,
		sb:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

// Get returns the deck with the given id.
func (r *Repository) Get(ctx context.Context, id string) (*deck.Deck, error) {
	saved := Deck{
		ID: id,
	}

	sql, args, err := r.sb.Select("id", "shuffled", "composition", "codes").
		From(table).
		Where(sq.Eq{"id": id}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build select sql for deck '%s': %w", id, err)
	}

	if err = r.pgx.QueryRow(ctx, sql, args...).
		Scan(&saved.ID, &saved.Shuffled, &saved.Composition, &saved.Codes); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("query row for deck '%s': %w", id, errs.ErrDeckNotFound)
		}

		return nil, fmt.Errorf("query row for deck '%s': %w", id, err)
	}

	d, err := DeckToDomainDeck(&saved)
	if err != nil {
		return nil, fmt.Errorf("deck '%s': %w", id, err)
	}

	r.log.Debugf("get deck '%s'", d.ID())

	return d, nil
}

// Save saves the given deck.
func (r *Repository) Save(ctx context.Context, d *deck.Deck) error {
	save := DomainDeckToDeck(d)

	sql, args, err := r.sb.Insert(table).
		Columns("id", "shuffled", "composition", "codes").
		Values(save.ID, save.Shuffled, save.Composition, save.Codes).
		Suffix(`ON CONFLICT (id) DO 
		UPDATE SET 
			shuffled = EXCLUDED.shuffled, 
			codes = EXCLUDED.codes, 
			updated_at = NOW()`).
		ToSql()
	if err != nil {
		return fmt.Errorf("build upsert sql for deck '%s': %w", d.ID(), err)
	}

	if _, err := r.pgx.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("save deck '%s': %w", d.ID(), err)
	}

	r.log.Debugf("saved deck '%s'", save.ID)

	return nil
}

// Migrate migrates the deck model.
func (r *Repository) Migrate(_ context.Context) error {
	d, err := iofs.New(migrations, "migrations")
	if err != nil {
		return fmt.Errorf("iofs: %w", err)
	}

	driver, err := mpgx.WithInstance(r.db, &mpgx.Config{})
	if err != nil {
		return fmt.Errorf("retrieve database.Driver instance: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		return fmt.Errorf("new migrate instance: %w", err)
	}

	if err := m.Migrate(version); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate up: %w", err)
	}

	r.log.Debug("migrated deck model")

	return nil
}

// Close closes any external connection in the repository.
func (r *Repository) Close(ctx context.Context) error {
	return r.pgx.Close(ctx)
}
