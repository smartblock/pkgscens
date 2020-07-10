# ENS with Go

Go library for Email, Notification, SMS

<hr />

## Installation

```ubuntu
$ go get github.com/smartblock/pkgscens@1.0.0
```

<hr />

## Usage
```go
package main

import (
	"fmt"

	"github.com/smartblock/pkgscens"
)

func main() {
	smtpAuth := pkgscens.SMTPAuth{
		Identity: "",
		Username: "example@gmail.com",
		Password: "your_password",
		Host:     "smtp.gmail.com",
	}

	sendMailInput := pkgscens.SendMailInput{
		Addr:     "smtp.gmail.com:587",
		SMTPAuth: smtpAuth,
		FromName: "Noreply Example",
		FromMail: "example@gmail.com",
		ToMail:   []string{"receipient@gmail.com"},
		ToName:   []string{"ReceipientName"},
    Subject:  "First Test Email",
    //HTML or TEXT
		MsgType:  "HTML",
		Message:  "<h1>My First Message</h1>",
	}

	err := pkgscens.SendMail(sendMailInput)

	if err != nil {
    //Error Return
		fmt.Println(err)
	}

  //Send Success
	fmt.Println("send success")
}
```
