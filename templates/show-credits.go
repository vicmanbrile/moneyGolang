package templates

var ShowCredits string = `

	<!DOCTYPE html>

	<head>
	
	<link rel="stylesheet" href="{{.StyleRecurse}}">
	<link rel="icon" type="image/x-icon" href="http://assets.localhost:8080/favicon.ico">
	
	</head>

	<body>

		<main>

			<table>

			<tr>
				<th>Name</th>
				<th>Price</th>
				<th>Paid</th>
				<th>Subtrac</th>
			</tr>
		
			{{range .Credits}}
				<tr>
					<td>{{ .Name }}</td>
					<td>{{ .Price }}</td>
					<td>{{ .Paid }}</td>
					<td>{{ .Subtrac }}</td>
				</tr>
			{{end}}

			</table>
	
		</main>
	
	</body>

	<html>
`
