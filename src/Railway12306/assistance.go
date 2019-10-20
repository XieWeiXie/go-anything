package Railway12306

func codeForStations(value string) string {
	if v, ok := ChPlace[value]; ok {
		return v
	}
	return ""

}
