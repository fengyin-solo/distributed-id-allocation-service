package leaseeventcollect

import (
	"context"
	"idgenerator/leaseeventsource"
)

func Collect(ctx context.Context, items []string, failAt int) ([]string, error) {
	data, _ := leaseeventsource.Stream(items, failAt)
	out := make([]string, 0, len(items))
	for item := range data {
		out = append(out, item)
	}
	return out, nil
}
