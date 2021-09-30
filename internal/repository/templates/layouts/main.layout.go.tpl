{{define "base"}}
    <!doctype html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport"
              content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Document</title>
        <link rel="stylesheet" href="/static/css/bootstrap.min.css">
        <link rel="stylesheet" href="/static/css/custom-styles.css">
        {{block "css" . }}
        {{end}}
    </head>
    <body>
    {{template "header" . }}

    {{block "content" . }}

    {{end}}
    

    {{template "footer" . }}

        {{block "js" . }}
        {{end}}
        <script src="/static/js/bootstrap.bundle.min.js"></script>
    </body>
    </html>
{{end}}