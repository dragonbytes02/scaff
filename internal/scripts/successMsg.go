package scripts

func SuccessMsg(selection ScriptApp) string {
	switch selection {
	case JsFuntionalClean | JsFuntionalHexa | JsFuntionalMcs | JsFuntionalMC | JsPooClean | JsPooHexa | JsPooMcs | JsPooMC:
		return "install success!\nexecute\npnpm i"
	}
	return ""
}
