package util

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"io"
	"strconv"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"gopkg.in/gomail.v2"
)

func SendMail(
	to string,
	subject string,
	text string,
	html string,
) error {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.SMTP_FROM)
	mailer.SetHeader("To", to)
	mailer.SetAddressHeader("Cc", config.SMTP_BCC, "noReply")
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", html)
	mailer.AddAlternative("text/plain", text)

	smtpPort, err := strconv.Atoi(config.SMTP_PORT)
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(
		config.SMTP_HOST,
		smtpPort,
		config.SMTP_USERNAME,
		config.SMTP_PASSWORD,
	)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

func CreateVerificationUrl(path string, id string, email string) string {
	token := CreateSha1(email)
	expires := TimeNowUnixEpoch() + config.EMAIL_VERIFICATION_TIMEOUT
	url := fmt.Sprintf(
		"%v%v?id=%v&token=%v&expires=%v",
		config.APP_CLIENT_URL,
		path,
		id,
		token,
		expires,
	)
	// `${APP_CLIENT_URL}/email-verify?id=${data._id}&token=${token}&expires=${expires}`;
	url = joinSignature(url)
	return url
}

func joinSignature(url string) string {
	signature := CreateSha256(url)
	// `${APP_CLIENT_URL}/email-verify?id=${data._id}&token=${token}&expires=${expires}&signature=${signature}`;
	return fmt.Sprintf("%v&signature=%v", url, signature)
}

func IsValidVerificationUrl(
	path string,
	id string,
	token string,
	expires string,
	signature string,
	requestURL string,
) bool {
	url := fmt.Sprintf(
		"%v%v?id=%v&token=%v&expires=%v&signature=%v",
		config.APP_CLIENT_URL,
		path,
		id,
		token,
		expires,
		signature,
	)

	if requestURL != url {
		return false
	}
	now := TimeNowUnixEpoch()
	expireInt64, err := strconv.ParseInt(expires, 10, 64)
	if err != nil {
		return false
	}
	if expireInt64 < now {
		return false
	}
	return true
}

func GenerateRandomBytes() ([]byte, error) {
	b := make([]byte, config.PASSWORD_RESET_BYTES)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func CreateSha1(in string) string {
	h := sha1.New()
	io.WriteString(h, in)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func CreateSha256(in string) string {
	h := sha256.New()
	h.Write([]byte(config.APP_SECRET))
	h.Write([]byte(in))
	return fmt.Sprintf("%x", h.Sum(nil))
}
