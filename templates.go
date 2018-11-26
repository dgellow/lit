package lit

import "text/template"

var (
	sideBySideTempl = template.Must(template.New("").Parse(`
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="utf-8">
				<style>
					body {
						margin: 0;
						padding: 0;
					}

					h1 {
						padding-left: 40px;
						padding-right: 40px;
						text-align: center;
						box-sizing: border-box;
						font-family: "Copperplate", serif;
						font-weight: 100;
					}
					@media (min-width: 900px) {
						h1 {width: 44%}
					}

					#content {
						max-width: 1600px;
						margin: auto;
						padding-top: 40px;
					}

					.section {
						display: flex;
						flex-direction: column;
					}
					@media (min-width: 900px) {
						.section {flex-direction: row}
					}

					.section .comments, .section .code {
						padding-left: 20px;
						padding-right: 20px;
						padding-top: 40px;
						padding-bottom: 40px;
						box-sizing: border-box;
						margin: 0;
					}
					.section .comments {
						padding-right: 80px;
						padding-left: 80px;
						border-bottom: 1px solid #f2f2f2;
						font-family: "Helvetica Neue", sans-serif;
						text-align: justify;
					}
					.section .code {
						color: #dedede;
						background-color: #122;
						font-family: "Menlo", monospace;
						overflow-x: scroll;
					}
					@media (min-width: 900px) {
						.section .comments {width: 44%}
						.section .code {width: 55%}
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
					<h1>{{.Filename}}</h1>
					<p></p>
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
