package scripts

func SuccessMsg(selection ScriptApp) string {
	switch selection {
	case JsFuntionalClean | JsFuntionalHexa | JsFuntionalMcs | JsFuntionalMC | JsPooClean | JsPooHexa | JsPooMcs | JsPooMC:
		return "\ninstall success!\nexecute\npnpm i"
	}
	return ""
}
