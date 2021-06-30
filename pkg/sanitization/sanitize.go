/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package sanitization

import "strings"

func SanitizeEmail(e string) string {
	trimmed := strings.ToLower(strings.TrimSpace(e))
	length := len(trimmed)

	if length < maxEmailLength {
		return trimmed
	}

	return trimmed[:maxEmailLength]
}

func SanitizeName(p string) string {
	trimmed := strings.TrimSpace(p)
	length := len(trimmed)

	if length < maxNameLength {
		return trimmed
	}

	return trimmed[:maxNameLength]
}

func SanitizePassword(p string) string {
	length := len(p)

	if length < maxPasswordLength {
		return p
	}

	return p[:maxPasswordLength]
}

func SanitizeToken(t string) string {
	trimmed := strings.TrimSpace(t)
	length := len(trimmed)

	if length < maxTokenLength {
		return trimmed
	}

	return trimmed[:maxTokenLength]
}
