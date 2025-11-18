package scripts

func SuccessMsg(selection ScriptApp) string {
	switch selection {
	case JsFuntionalClean, JsFuntionalHexa, JsFuntionalMcs, JsFuntionalMC, JsPooClean, JsPooHexa, JsPooMcs, JsPooMC:
		return "\nconfigure your environment.\ninstall nodeJs, Npm and Pnpm\nexecute:\n\npnpm i\n"

	case GoHexa, GoMcs, GoMC:
		return "\nconfigure your environment.\nexecute:\n\ngo mod tidy\n"
	}

	return ""
}
