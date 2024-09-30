package handlers

import "strings"

func GenerateSlug(name string) string {

	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

// Helper function to generate a thumb URL based on the product type name
func GenerateThumb(name string) string {

	return "/images/thumbs/" + strings.ToLower(strings.ReplaceAll(name, " ", "-")) + ".jpg"
}
