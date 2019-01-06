package email

import (
  "html/template"
  "bytes"

  "github.com/cagox/gge/config"
)


//ParseHTMLEmail takes the given information and returns a prepared email body.
func ParseHTMLEmail(subtemplate string, data interface{}) (string, error) {
  basetemplate := config.Config.TemplateRoot+"/email/base.html"
  body := ""

  t := template.New("base.html")
  t, err := t.ParseFiles(basetemplate, subtemplate)
  if err != nil {
    return "", err
  }

  buffer := new(bytes.Buffer)

  if err = t.Execute(buffer, data); err != nil {
    return "", nil
  }

  body = buffer.String()

  return body, nil
}
