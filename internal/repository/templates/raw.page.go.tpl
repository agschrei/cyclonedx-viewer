{{template "base" .}}

{{define "css"}}
 <link rel="stylesheet" href="/static/css/monokai-sublime.min.css">
{{end}}

{{define "content"}}

<div class="container">
        <div class="row my-5">
        <h1 class="mb-4">Raw Input</h1>
          <pre class="px-5"><code class="mb-4 language-json rounded-corners">{{ . }}</code></pre>
        </div>
</div>

{{end}}

{{define "js"}}
<script src="/static/js/highlight.min.js"></script>
<script src="/static/js/highlightjs-line-numbers.min.js"></script>
<script>
    hljs.highlightAll();
    hljs.initLineNumbersOnLoad();
</script>
{{end}}