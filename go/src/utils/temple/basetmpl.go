package temple

// baseTmpl is the default page template rendered by NewPage
var baseTmpl = `
{{define "base"}}
<!doctype html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title>{{.Title}}</title>
        {{range $_, $tag := .MetaTags }}
            <meta name="{{$tag.Name}}" content="{{$tag.Content}}">
        {{end}}

        {{range $_, $css := .CSS }}
            <link rel="stylesheet" href="{{$css}}">
        {{end}}

        {{range $_, $js := .JS }}
            <script src="{{$js}}"></script>
        {{end}}

    </head>
    <body>
    <div class="container">
        {{range $_, $child := .Children}}
        <div class="$child.Class">
                {{renderChild($child)}}
        </div>
        {{end}}
    </div>
    </body>
</html>
{{end}}
`
