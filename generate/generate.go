package generate

type Render interface {
	// Single parse the single post content, return the config map and html content bytes
	Single(raw []byte) (map[string]string, []byte, error)
}
