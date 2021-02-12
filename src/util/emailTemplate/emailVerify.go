package emailTemplate

import (
	"bytes"
	"html/template"
)

func EmailVerify(url string) (string, error) {
	template, err := template.ParseFiles("./template/email-verify.html")
	if err != nil {
		return "", err
	}
	var data = map[string]interface{}{
		"url": url,
	}
	var templateByte bytes.Buffer
	template.Execute(&templateByte, data)

	return templateByte.String(), nil
}
