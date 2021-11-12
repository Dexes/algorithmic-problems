package platform

func getLanguageTitle(lang string) string {
	switch lang {
	case "cpp":
		return "C++"
	case "csharp":
		return "C#"
	case "sql":
		return "SQL"
	default:
		return lang
	}
}

func getLanguageExtension(lang string) string {
	switch lang {
	case "python":
		return "py"
	case "csharp":
		return "cs"
	default:
		return lang
	}
}
