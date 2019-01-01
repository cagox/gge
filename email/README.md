# github.com/cagox/gge/email



## Email Data struct

```
type Data struct {
  To         string
  Name       string //Optional
  Subject    string
  Body       string
  HTML       bool  //Optional. Defaults to false.


```

This struct holds all the data that is passed to the SystemEmail method. To is
the email address of the individual that the email is going to.

**To**: is the email address that you are sending the message to.

**Name**: is the "proper name" of the person the email is being sent to. This
will often be blank, but I am including it for future flexibility.

**Subject**: is the subject of the email.

**Body**: is the body of the email, as the sender wants it to be sent. This
will be a processed string of some sort, or the .String() output of a processed
template. The mailer isn't going to make any further changes to this part and
will send it as is.

**HTML**: is a boolean value. If it is set to true, the mailer will treat the
email as an HTML email. Otherwise it will treat it as plain text. The main
difference is a content header that is set to let the email client know what it
is dealing with.



## email.SystemEmail(data Data)

```
func SystemEmail(data Data) {...}
```

This is the mailer.  While I may add more detail to this section later, the
important parts are basically laid out above in the description of the Data
struct.

It is important to note, however, that in order for it to work properly, some
values have to be set in the GGE configuration file.

```
"FromAddress"       : "cagox@cagox.net",
"FromName"          : "Cagox Media",
"SMTPServer"        : "email.server.com:465",
"SMTPUserName"      : "someone@someplace.net",
"SMTPPassword"       : "password"
```

These are the items used to describe who the email is from, and to handle
authentication with the SMTP server.

At a future time, I may add options to allow for connections other than ssl, or
the ability to use senmail or similar on localhost.
