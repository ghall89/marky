package files

import "fmt"

func FormatHTML(html []byte, title string) []byte {
	output := fmt.Sprintf(`
  <!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="nano.css">
    <title>%s</title>
  </head>
  <body>
    <main>
    %s
    </main>
  </body>
  </html>
  `, title, string(html))

	return []byte(output)
}
