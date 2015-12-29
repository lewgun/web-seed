package misc
import "regexp"

var (
	reName  *regexp.Regexp
	reEmail *regexp.Regexp
	rePhone *regexp.Regexp
)

func init() {
	reName = regexp.MustCompile(`^[a-zA-Z_]+[0-9]*$`)
	reEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	rePhone = regexp.MustCompile(`^(0|\+86|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$`)
}


func ValidUserName(name string) bool {

	return reName.MatchString(name)
}

func ValidEmail(email string) bool {
	return reEmail.MatchString(email)
}

func ValidPhone(phone string) bool {
	return rePhone.MatchString(phone)
}
