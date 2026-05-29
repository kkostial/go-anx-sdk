package paging

import (
	"context"
	"iter"
)

// PagedResponse is a generic response type wrapping paging logic of the anexia engine.
type PagedResponse[T any] struct {
	Page       int `json:"page"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
	Limit      int `json:"limit"`
	Data       []T `json:"data"`
}

// PageFetcher is a function that fetches the desired page from the api.
type PageFetcher[T any] func(
	ctx context.Context,
	page int,
	limit int,
) (PagedResponse[T], error)

const engineMaxPageLimit = 100

// Paginate iterates all resources from a paged endpoint using the provided PageFetcher.
func Paginate[T any]( //revive:disable:cognitive-complexity
	ctx context.Context,
	fetchPage PageFetcher[T],
) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var zero T
		page := 1

		for {
			err := ctx.Err()
			if err != nil {
				yield(zero, err)
				return
			}

			resp, err := fetchPage(ctx, page, engineMaxPageLimit)
			if err != nil {
				yield(zero, err)
				return
			}

			for _, v := range resp.Data {
				if !yield(v, nil) {
					return
				}
			}

			if page >= resp.TotalPages {
				return
			}

			page++
		}
	}
}
