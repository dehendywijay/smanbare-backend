package utility

import (
    "regexp"
    "strings"

)


func MakeSlug(title string) string {
    slug := strings.ToLower(title)

    reg := regexp.MustCompile(`[^a-z0-9]+`)
    slug = reg.ReplaceAllString(slug, "-")

    slug = strings.Trim(slug, "-")

    return slug
}

func SlugToTitle(slug string) string {
    // ganti "-" jadi spasi
    title := strings.ReplaceAll(slug, "-", " ")

    // kapitalisasi tiap kata
    words := strings.Fields(title)
    for i, word := range words {
        words[i] = strings.Title(word)
    }

    return strings.Join(words, " ")
}