package templates

var FormCredit string = `
	{{if .Success}}
		<h1>Thanks for your message!</h1>
	{{else}}
		<h1>Compra a Credito</h1>
		<form method="POST">
			<label>Nombre:</label><br />
			<input type="text"><br />
			<label>Precio:</label><br />
			<input type="text"><br />
			<label>Message:</label><br />
			<textarea name="message"></textarea><br />
			<input type="submit">
		</form>
	{{end}}
`
