package business

import "context"

type Business struct {
	Context context.Context
}

func NewBusiness(ctx context.Context) *Business {
	return &Business{
		Context: ctx,
	}
}
