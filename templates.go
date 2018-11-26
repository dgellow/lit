package lit

import "text/template"

var (
	sideBySideTempl = template.Must(template.New("").Parse(`
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="utf-8">
				<style>
					.section .comments, .section .code {
						display: inline-block;
						padding: 20px;
						box-sizing: border-box;
						margin: 0;
					}
					.section .comments {
						width: 44%;
						border-bottom: 1px solid grey;
					}
					.section .code {
						width: 55%;
						color: #dedede;
						background-color: #122;
						font-family: monospace;
					}
					.chunk {
						display: inline;
					}
					.is-keyword {
						color: #75c0de;
					}
					.is-literal {}
					.is-operator {
						color: white;
					}
					.string {
						color: #d8a66d;
					}
					.func {
					  color: #ff026f;
					}
				</style>
			</head>
			<body>
				<div id="content">
					{{range .Sections}}
					<div class="section">
						<div class="comments">{{.Comments}}</div>
						<div class="code"><pre><code>{{.Code}}</code></pre></div>
					</div>
					{{end}}
				</div>
			</body>
		</html>
	`))
)
