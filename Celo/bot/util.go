package bot

func boldText(str string) string {
    return "*" + str + "*"
}

func warnText(str string) string {
    return "\xE2\x9A\xA0 " + str
}

func errText(str string) string {
    return "\xE2\x9D\x8C " + str
}