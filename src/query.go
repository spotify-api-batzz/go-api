package main

import (
	"fmt"
	"net/url"
	"strconv"
)

type QueryOptions struct {
	Paginate int `json:"paginate"`
	Page     int `json:"page"`
}

func defaultQueryOptions() *QueryOptions {
	return &QueryOptions{
		Paginate: 0,
		Page:     0,
	}
}

func Transform(query string, opts *QueryOptions) string {
	if opts.Page != 0 && opts.Paginate != 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, opts.Paginate)
	}
	if opts.Paginate != 0 {
		query = fmt.Sprintf("%s OFFSET %d", query, opts.Page*opts.Paginate)
	}
	return query
}

func parseQueryOptions(vars url.Values) (*QueryOptions, error) {
	options := defaultQueryOptions()
	availableVars := []string{"paginate", "page"}
	for _, key := range availableVars {
		value := vars.Get(key)
		if value == "" {
			continue
		}
		switch key {
		case "paginate":
			intVal, err := strconv.Atoi(value)
			if err != nil {
				return options, err
			}
			options.Paginate = intVal
			continue
		case "page":
			intVal, err := strconv.Atoi(value)
			if err != nil {
				return options, err
			}
			options.Page = intVal
			continue
		}
	}
	return options, nil
}
