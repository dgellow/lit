package lit

import "text/template"

var (
	// DocumentTempl is used when generating a full HTML document.
	DocumentTempl = template.Must(template.New("").Parse(`
<! DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<style>body {margin: 0; padding: 0}</style>
	</head>
	<body>
		{{printf "%s" .}}
	</body>
</html>
`))

	// SideBySideTempl generates a side by side page, with comments on the left
	// and code on the right. If the document width goes below 900px sections
	// are instead displayed in one unique column, with comments, then code.
	SideBySideTempl = template.Must(template.New("").Parse(`
<style>
	#content-lit h1 {
		padding-left: 40px;
		padding-right: 40px;
		padding-top: 20px;
		padding-bottom: 20px;
		text-align: center;
		box-sizing: border-box;
		font-family: "Copperplate", serif;
		font-weight: 100;
		position: sticky;
		top: 0;
		margin: 0;
		background-color: white;
		box-shadow: 0 6px 12px -13px black;
	}
	@media (min-width: 900px) {
		#content-lit h1 {width: 44%}
	}

	#content-lit {
		max-width: 1600px;
		margin: auto;
		padding-top: 40px;
	}

	#content-lit .section {
		display: flex;
		flex-direction: column;
	}
	@media (min-width: 900px) {
		#content-lit .section {flex-direction: row}
	}

	#content-lit .section .comments,
	#content-lit .section .code {
		padding-left: 20px;
		padding-right: 20px;
		padding-top: 40px;
		padding-bottom: 40px;
		box-sizing: border-box;
		margin: 0;
	}
	#content-lit .section .comments {
		padding-right: 80px;
		padding-left: 80px;
		border-bottom: 1px solid #f2f2f2;
		font-family: "Helvetica Neue", sans-serif;
		text-align: justify;
	}
	#content-lit .section .code {
		color: #dedede;
		background-color: #122;
		font-family: "Menlo", monospace;
		overflow-x: scroll;
	}
	@media (min-width: 900px) {
		#content-lit .section .comments {width: 44%}
		#content-lit .section .code {width: 55%}
	}

	#content-lit .chunk {
		display: inline;
	}
	#content-lit .is-keyword {
		color: #75c0de;
	}
	#content-lit .is-literal {}
	#content-lit .is-operator {
		color: white;
	}
	#content-lit .string {
		color: #d8a66d;
	}
	#content-lit .func {
		color: #ff026f;
	}
</style>
<div id="content-lit">
	<h1>{{.Filename}}</h1>
	{{range .Sections}}
	<div class="section">
		<div class="comments">{{.Comments}}</div>
		<div class="code">
			<pre><code>{{.Code}}</code></pre>
		</div>
	</div>
	{{end}}
</div>`))
)
