package main

const yayMessageInEnglish = "yay from func, "
const yayMessageInSpanish = "Hola que tal, "

func Yay(name string, lang string) string {
	return yayPrefix(lang) + name
}

func yayPrefix(lang string) (prefix string) {
	switch lang {
	case "ENG":
		prefix = yayMessageInEnglish
	case "SPA":
		prefix = yayMessageInSpanish
	default:
		prefix = yayMessageInEnglish
	}

	return
}
