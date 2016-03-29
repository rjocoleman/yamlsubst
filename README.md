# YAMLsubst

Takes YAML from a specified file and substitutes into a template from `stdin`, output is sent to `stdout`.

Templates are written in Go's [`text/template`](http://golang.org/pkg/text/template/).

### Example Usage:

`yamlsubst -yaml nginx.conf.yml < nginx.conf.template > /etc/nginx/nginx.conf`

## Template Functions

### datetime

Alias for time.Now

```
# Generated at {{datetime}}
```

Outputs:

```
# Generated at 2015-01-23 13:34:56.093250283 -0800 PST
```

```
# Generated at {datetime.Format("Jan 2, 2006 at 3:04pm (MST)")}
```

Outputs:

```
# Generated at Jan 23, 2015 at 1:34pm (EST)
```

See the time package for more usage: http://golang.org/pkg/time/

### split

Wrapper for [strings.Split](http://golang.org/pkg/strings/#Split). Splits the input string on the separating string and returns a slice of substrings.

```
{{ $url := .StringToSplit ":" }}
    host: {{index $url 0}}
    port: {{index $url 1}}
```

### toUpper

Alias for [strings.ToUpper](http://golang.org/pkg/strings/#ToUpper) Returns uppercased string.

```
key: {{toUpper "value"}}
```

### toLower

Alias for [strings.ToLower](http://golang.org/pkg/strings/#ToLower). Returns lowercased string.

```
key: {{toLower "Value"}}
```

### join

Alias for the [strings.Join](http://golang.org/pkg/strings/#Join) function to work with YAML arrays.

```
{{join .Array ","}}
```

### replace

Alias for the [strings.Replace](http://golang.org/pkg/strings/#Replace) function.

```
{{replace .replaceExample "-" "_" -1}}
```

## Example Usage

Located at `./test.yml`

```yaml
---
place: World
foods:
  - apples
  - bananas
  - avocados
  - pears
  - carrots

items: foo,bar,baz
replaceExample: qux-bar-bat

```

Located at `./test.template`

```text
Hello {{ .place }} it's {{datetime}}.

My favorite foods are {{join .foods ", "}}.

{{$items := split .items ","}}
{{range $index, $element := $items}}{{$index}}: {{$element}}
{{end}}

Upper: {{toUpper .place}}

Lower: {{toLower .place}}

{{replace .replaceExample "-" "_" -1}}
```

Usage: `yamlsubst -yaml test.yml < test.template `

Output:

```text
Hello World it's 2016-03-29 18:28:05.793361229 +1300 NZDT.

My favorite foods are apples, bananas, avocados, pears, carrots.


0: foo
1: bar
2: baz


Upper: WORLD

Lower: world

qux_bar_bat
```

Go's [`text/template`](http://golang.org/pkg/text/template/) package is very powerful. For more details on it's capabilities see its [documentation.](http://golang.org/pkg/text/template/)


### See Also

Inspired by:

- https://github.com/kelseyhightower/confd
- https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html
