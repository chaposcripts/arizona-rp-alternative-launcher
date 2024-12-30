package main

import (
	"context"

	sampquery "github.com/Southclaws/go-samp-query"
)

func getServerQueryInfo(host string, port int) (sampquery.Server, error) {
	var result sampquery.Server
	query, err := sampquery.NewQuery(host)
	if err != nil {
		return result, err
	}
	var ctx context.Context
	result, err = query.GetInfo(ctx, true)
	if err != nil {
		return result, err
	}
	return result, nil
}
