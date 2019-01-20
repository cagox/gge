package email

import (
    //"fmt"
	  //"log"
    //"net"
    //"net/mail"
	  //"net/smtp"
    //"crypto/tls"

    "gopkg.in/mail.v2"

    "github.com/cagox/gge/config"

)



//Data provides information for the SystemEmail function to do it's job.
type Data struct {
  To         string
  Name       string //Optional
  Subject    string
  Body       string
  HTML       bool  //Optional. Defaults to false.
}


func SystemEmail(data Data) error {
  m := mail.NewMessage()
  m.SetHeader("From", config.Config.FromAddress)
  m.SetHeader("To", data.To)
  //m.SetAddres("Cc", "some@address.net", "Name")
  m.SetHeader("Subject", data.Subject)
  if data.HTML {
    m.SetBody("text/html", data.Body)
  } else {
    m.SetBody("text/plain", data.Body)
  }
  //m.Attach("/some/attached/file")

  d := mail.NewDialer(config.Config.SMTPServer, config.Config.SMTPPort, config.Config.SMTPUserName, config.Config.SMTPPassword)
  d.StartTLSPolicy = mail.MandatoryStartTLS

  //Now send the mail.

  if err := d.DialAndSend(m); err != nil {
    return err
  }

  return nil
}

/*

//SystemEmail lets the system send an email out to an individual.
func SystemEmail(data Data) error {
  from := mail.Address{Name: config.Config.FromName, Address: config.Config.FromAddress}
  to := mail.Address{Address: data.To}
  if data.Name != "" {
    to.Name = data.Name
  }
  subj := data.Subject

  headers := make(map[string]string)
  headers["From"] = from.String()
  headers["To"] = to.String()
  headers["Subject"] = subj
  if data.HTML {
    headers["MIME-version"] = "1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
  }


  // Setup message
  message := ""
  for k,v := range headers {
      message += fmt.Sprintf("%s: %s\r\n", k, v)
  }
  message += "\r\n" + data.Body

  // Connect to the SMTP Server
  servername := config.Config.SMTPServer

  host, _, _ := net.SplitHostPort(servername)

  //smtp.PlainAuth("","cagox@cagox.net", "!!23ABcd1", host)
  auth := smtp.PlainAuth("",config.Config.SMTPUserName, config.Config.SMTPPassword, host)

  // TLS config
  tlsconfig := &tls.Config {
      InsecureSkipVerify: true,
      ServerName: host,
  }

  // Here is the key, you need to call tls.Dial instead of smtp.Dial
  // for smtp servers running on 465 that require an ssl connection
  // from the very beginning (no starttls)
  conn, err := tls.Dial("tcp", servername, tlsconfig)
  if err != nil {
      return err
  }

  c, err := smtp.NewClient(conn, host)
  if err != nil {
      return err
  }

  // Auth
  if err = c.Auth(auth); err != nil {
      return err
  }

  // To && From
  if err = c.Mail(from.Address); err != nil {
      return err
  }

  if err = c.Rcpt(to.Address); err != nil {
      return err
  }

  // Data
  w, err := c.Data()
  if err != nil {
      return err
  }

  _, err = w.Write([]byte(message))
  if err != nil {
      return err
  }

  err = w.Close()
  if err != nil {
      return err
  }

  c.Quit()

  //fmt.Println(message)
  return nil
}
*/

/*
Test tests email.  It is literally a copy and paste of the gist that I got the info from.
https://gist.github.com/chrisgillis/10888032

I will leave it here, for now, for future reference.



func Test() {

    from := mail.Address{Name: "burlingk", Address: "cagox@cagox.net"}
    to   := mail.Address{Name: "burlingk", Address: "burlingk@cagox.net"}
    subj := "This is the email subject"
    body := "This is an example body.\n With two lines."

    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subj

    // Setup message
    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    // Connect to the SMTP Server
    servername := "secure.emailsrvr.com:465"

    host, _, _ := net.SplitHostPort(servername)

    auth := smtp.PlainAuth("","cagox@cagox.net", "!!23ABcd1", host)

    // TLS config
    tlsconfig := &tls.Config {
        InsecureSkipVerify: true,
        ServerName: host,
    }

    // Here is the key, you need to call tls.Dial instead of smtp.Dial
    // for smtp servers running on 465 that require an ssl connection
    // from the very beginning (no starttls)
    conn, err := tls.Dial("tcp", servername, tlsconfig)
    if err != nil {
        log.Panic(err)
    }

    c, err := smtp.NewClient(conn, host)
    if err != nil {
        log.Panic(err)
    }

    // Auth
    if err = c.Auth(auth); err != nil {
        log.Panic(err)
    }

    // To && From
    if err = c.Mail(from.Address); err != nil {
        log.Panic(err)
    }

    if err = c.Rcpt(to.Address); err != nil {
        log.Panic(err)
    }

    // Data
    w, err := c.Data()
    if err != nil {
        log.Panic(err)
    }

    _, err = w.Write([]byte(message))
    if err != nil {
        log.Panic(err)
    }

    err = w.Close()
    if err != nil {
        log.Panic(err)
    }

    c.Quit()

}

*/
