package controllers

import "net/http"

var graphiql = []byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.css"/>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/2.0.3/fetch.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/16.2.0/umd/react.production.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/16.2.0/umd/react-dom.production.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.js"></script>
		</head>
		<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
			<div id="graphiql" style="height: 100vh;">Loading...</div>
			<script>
				function fetchGQL(params) {
					return fetch("/graphql", {
						method: "post",
						body: JSON.stringify(params),
						credentials: "include",
					}).then(function (resp) {
						return resp.text();
					}).then(function (body) {
						try {
							return JSON.parse(body);
						} catch (error) {
							return body;
						}
					});
				}
				ReactDOM.render(
					React.createElement(GraphiQL, {fetcher: fetchGQL}),
					document.getElementById("graphiql")
				)
			</script>
		</body>
	</html>
	`)

// GraphiQLController handles requests that involves GraphiQL IDE
type GraphiQLController struct{}

// NewGraphiQLController creates new GraphiQLController
func NewGraphiQLController() *GraphiQLController {
	return &GraphiQLController{}
}

func (c GraphiQLController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		respond(res, jsonErr("Method not supported"), http.StatusMethodNotAllowed)
		return
	}

	res.Write(graphiql)
}
