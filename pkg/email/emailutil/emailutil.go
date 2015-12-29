package emailutil

import (
	"bytes"
	"text/template"

	"github.com/lewgun/web-seed/pkg/email"
	"github.com/lewgun/web-seed/pkg/errutil"
	"github.com/lewgun/web-seed/pkg/misc"
)

const (
	restTmpl = `
  Hi, {{if .Name}}{{.Name}}{{else}}Dear{{end}}:

  我们的系统收到一个请求，说您希望通过电子邮件重新设置您在【xx】的密码。您可以点击下面的链接重设密码：
	{{.Callback}}/resetpassword.html?mt={{.Type}}&ck={{.Code}}&me={{base64 .To}}


  如果以上链接无法点击，请将上面的地址复制到你的浏览器(如IE)的地址栏中打开,
  如果这个请求不是由您发起的，那没问题，您不用担心，您可以安全地忽略这封邮件!

      (这是一封自动产生的email，请勿回复.)`
)

//MailType 邮件类型
type MailType string

const (
	Reset  = "reset"  //重置密码邮件
	Unlock = "unlock" //帐户激活邮件
)

const (
	titleReset = "【xxx】重置密码"
)

const (
	KeyMailType = "mt"
	KeyCode     = "ck"
	KeyMedia    = "me"
)

const (
	defCallback = "callback url"
)

var t *template.Template

func init() {

	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"base64": misc.Base64,
	}

	t = template.Must(template.New("letter").Funcs(funcMap).Parse(restTmpl))
}

type Config struct {
	Name     string
	Callback string
	Type     string
	Code     string
	To       string
}

func compose(p *Config) (subject string, body []byte) {
	if p == nil {
		return "", nil
	}
	if p.Callback == "" {
		p.Callback = defCallback
	}

	switch p.Type {
	case Reset:
		buf := &bytes.Buffer{}
		t.Execute(buf, p)
		return titleReset, buf.Bytes()

	default:
		//placeholder
	}

	return "", nil
}

//SendMail send mail by config.
func SendMail(c *Config) error {

	if c == nil {
		return errutil.ErrIllegalParam
	}

	subject, body := compose(c)
	if subject == "" {
		return nil
	}

	return email.Send(subject, body, []string{c.To})
}
