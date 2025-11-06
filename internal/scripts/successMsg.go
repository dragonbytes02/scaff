package scripts

func SuccessMsg(selection ScriptApp) string {
	switch selection {
	case JsFuntionalClean | JsFuntionalHexa:
		return "Ejecuta pnpm instala node e instala npm"
	}
	return ""
}
