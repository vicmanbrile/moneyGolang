package templates

var FormSession string = `
<!DOCTYPE html>

<head>

<link rel="icon" type="image/x-icon" href="http://assets.localhost:8080/favicon.ico">

</head>

<body>

	<main>
		<h1>Iniciar Sesion</h1>
		<form method="POST">
			<label for="user">Usuario:</label><br />
			<input id="user" name="user" type="text"><br />
			<label for="pass">Contraseña:</label><br />
			<input id="pass" name="user" type="text"><br />
			<input type="submit">
		</form>

	</main>

</body>

<html>
	
`
