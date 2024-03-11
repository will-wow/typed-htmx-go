package blog

import (
	"embed"
	"fmt"
	"strings"
)

func parseSlugs(fs embed.FS) (map[string]string, error) {
	posts, err := fs.ReadDir("posts")
	if err != nil {
		return nil, fmt.Errorf("reading posts directory: %w", err)
	}

	slugs := make(map[string]string, len(posts))

	for _, post := range posts {
		title := post.Name()
		slug := normalizeSlug(title)

		_, exists := slugs[slug]
		if exists {
			return nil, fmt.Errorf("duplicate slug: %s", slug)
		}

		slugs[slug] = title
	}

	return slugs, nil
}

// normalizeSlug normalizes a user provided slug to a slug that can be used to look up a recipe
func normalizeSlug(slug string) string {
	// Remove the extension
	slug = removeExtension(slug)
	return snakeCaseFilename(slug)
}

// removeExtension removes the .md extension from a filename
func removeExtension(filename string) string {
	return strings.Replace(filename, ".md", "", 1)
}

// snakeCaseFilename converts a filename to snake case
func snakeCaseFilename(filename string) string {
	// Replace spaces with underscores
	filename = strings.Replace(filename, " ", "_", -1)
	// Strip out commas
	filename = strings.Replace(filename, ",", "", -1)
	return strings.ToLower(filename)
}
