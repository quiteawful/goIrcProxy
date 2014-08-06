package main

var HtmlLogHead string = `<!DOCTYPE html><html><head><meta http-equiv="refresh" content="5"></head><body>`
var HtmlLogEnd string = "</body></html>"

var HtmlMain string = `<!DOCTYPE html>
<html><title>Doclol Irc Proxy</title></head><body>
<iframe src="./log" height="100px"><p>Deine Admins suck0rn.</p></iframe>
<form action="/" method="post">
Nachricht: <input type="text" name="content">
<input type="submit">
</form>
</body></html>
`
