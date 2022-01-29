package platform

import "strings"

func getLanguageTitle(lang string) string {
	switch lang {
	case "cpp":
		return "C++"
	case "csharp":
		return "C#"
	case "sql":
		return "SQL"
	default:
		return upperFirst(lang)
	}
}

func getLanguageExtension(lang string) string {
	switch lang {
	case "python":
		return "py"
	case "csharp":
		return "cs"
	case "bash":
		return "sh"
	default:
		return lang
	}
}

func pathEscape(s string) string {
	s = strings.ReplaceAll(s, "%", "%25")
	s = strings.ReplaceAll(s, " ", "%20")
	return s
}

func upperFirst(s string) string {
	if 'A' <= s[0] && s[0] <= 'Z' {
		return s
	}

	data := []byte(s)
	data[0] -= 32

	return string(data)
}
