package email_operator

import (
	"encoding/base64"
	"log"

	"github.com/go-gomail/gomail"
	"github.com/wuxiaoxiaoshen/go-anything/configs"
)

type (
	EmailAction struct {
		Sender   string      `json:"sender"`
		Receiver []string    `json:"receiver"`
		Subject  string      `json:"subject"`
		Content  interface{} `json:"content"`
		message  *gomail.Message
		dialer   *gomail.Dialer
	}
)

var DefaultEmailAction = &EmailAction{}

func EmailInit() {
	emails := configs.DefaultConfigs.LoadConfigs("email")
	var result map[string]interface{}
	result = emails.(map[string]interface{})
	sender := result["sender"].(string)
	receivers := result["receivers"].([]interface{})
	auth := result["auth"].(string)
	word, _ := base64.StdEncoding.DecodeString(auth)
	action := newEmailAction(sender, string(word))
	do := func(v []interface{}) []string {
		var result []string
		for _, i := range v {
			result = append(result, i.(string))
		}
		return result
	}
	action.Receiver = do(receivers)
	DefaultEmailAction = action
}

func newEmailAction(sender string, auth string) *EmailAction {
	return &EmailAction{
		Sender:  sender,
		message: gomail.NewMessage(),
		dialer:  gomail.NewDialer("smtp.qq.com", 465, sender, auth),
	}
}

func (E *EmailAction) AddReceiver(names []string) {
	E.Receiver = append(E.Receiver, names...)
}
func (E *EmailAction) AddContent(v interface{}) {
	E.Content = v
}

func (E *EmailAction) formatReceiver() []string {
	var result []string
	for _, e := range E.Receiver {
		result = append(result, E.message.FormatAddress(e, "收件人"))
	}
	return result
}
func (E *EmailAction) Run(subject string) bool {
	E.message.SetAddressHeader("From", E.Sender, "天空之城")
	E.message.SetHeader("To", E.formatReceiver()...)
	E.message.SetHeader("Subject", subject)
	E.message.SetBody("text/html", E.Content.(string))
	e := E.dialer.DialAndSend(E.message)
	if e != nil {
		log.Println("email: send email :", e.Error())
		return false
	}
	return true

}

func (E *EmailAction) Close() {

}
