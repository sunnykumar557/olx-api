package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/sunnykumar557/olx-api/internal/middlewares"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListingHandler struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewListingHandler(db *sql.DB, logger *slog.Logger) *ListingHandler {
	return &ListingHandler{
		db:     db,
		logger: logger,
	}
}

func (lh ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rows, err := lh.db.QueryContext(
		ctx,
		`SELECT id,title,description,price,city,created_at
			FROM listings
			ORDER BY created_at DESC
			LIMIT 100`)

	if err != nil {
		//log.Printf("query: %v", err)
		lh.logger.Error("listings query error", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	listings := []listing{}
	for rows.Next() {
		var l listing
		if err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt); err != nil {
			lh.logger.Error("rows scan error", "err", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		lh.logger.Info("listings fetched", "total", len(listings))

		listings = append(listings, l)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows.Err: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(listings)
}

func (lh ListingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//requestId := ctx.Value("requestCtxId").(string)
	requestId := middlewares.RequestIDFromContext(ctx)

	id := r.PathValue("id")
	fmt.Println(id)

	_, err := lh.db.ExecContext(ctx, `
		    DELETE FROM listing
			WHERE id = $1`, id)

	if err != nil {
		//log.Printf("db.delete: %v", err)

		lh.logger.Error("delete failed", "listing_id", id, "request_id", requestId, "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
