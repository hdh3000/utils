package sendjobmail

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/mail"
	"path"
	"strings"

	"google.golang.org/api/gmail/v1"
)

// TmplArgs are the set of arguments available to all templates.
type TmplArgs struct {
	To   *BasicInfo
	From *BasicInfo
}

type BasicInfo struct {
	Email       string
	First       string
	Last        string
	CompanyName string
	Title       string
}

func ParseTmpl(name string, tmplDir string, args *TmplArgs) (string, error) {
	path := path.Join(tmplDir, fmt.Sprintf("%s.tmpl", name))

	buf := &bytes.Buffer{}
	if err := template.Must(template.ParseFiles(path)).Execute(buf, args); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func SendEmail(ctx context.Context, gmailClient *gmail.Service, from, to *mail.Address, subject, body string) error {
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	var msg string
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body

	gmsg := gmail.Message{
		Raw: base64.RawURLEncoding.EncodeToString([]byte(msg)),
	}

	_, err := gmailClient.Users.Messages.Send("me", &gmsg).Do()

	return err
}

// HACKY AF (stack overflows finest)
func encodeRFC2047(s string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{s, ""}
	return strings.Trim(addr.String(), " <>")
}
