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
			<label for="user">Usuario:</label>
			<input id="user" name="user" type="text">
			<br />
			<label for="pass">Contraseña:</label>
			<input id="pass" name="pass" type="text"><br />
			<br />
			<input type="submit">
		</form>

	</main>

</body>

<html>
	
`
