package open

// URL opens a new browser window pointing to url.
func URL(url string) error {
	return open(url)
}
