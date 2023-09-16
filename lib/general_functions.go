package lib

import (
	"fmt"
	"net/url"
	"strings"
)

func buildPaginationQuery(pagination map[string]interface{}) (string, error) {
	validKeys := map[string]bool{
		"page":      true,
		"pageSize":  true,
		"name":      true,
		"use_regex": true,
	}
	for key := range pagination {
		if !validKeys[key] {
			return "", fmt.Errorf("invalid key '%s' in pagination", key)
		}
	}
	// Initialize an empty slice to hold the key-value pairs
	var queryParams []string

	// Iterate over the pagination map
	for key, value := range pagination {
		// Convert the value to a string
		valueStr := fmt.Sprintf("%v", value)

		// URL encode both the key and the value and join them with "="
		queryParam := fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(valueStr))

		// Append the query parameter to the slice
		queryParams = append(queryParams, queryParam)
	}

	// Join the query parameters with "&" to create the final query string
	queryString := strings.Join(queryParams, "&")

	// Append "pagination=true" to the query string
	queryString = "?" + queryString + "&pagination=true"

	return queryString, nil
}

// validateExchangeParams checks if required parameters are missing.
func validateExchangeParams(vhost, exchange, exchangeType string) error {
	if err := validateParam(vhost, "vhost"); err != nil {
		return err
	}
	if err := validateParam(exchange, "exchange"); err != nil {
		return err
	}
	if err := validateParam(exchangeType, "exchange type"); err != nil {
		return err
	}
	return nil
}

// validateQueueParams checks if required parameters are missing.
func validateQueueParams(vhost, queue string) error {
	if err := validateParam(vhost, "vhost"); err != nil {
		return err
	}
	if err := validateParam(queue, "queue"); err != nil {
		return err
	}
	return nil
}

func validateParam(param, paramName string) error {
	if param == "" {
		return fmt.Errorf("missing %s parameter", paramName)
	}
	return nil
}
