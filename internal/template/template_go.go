package template

const defaultGoTemplate = `/* Generated by protoc-gen-tpl. Do not edit! */

package {{ .GoPackage }}
{{- range $message := .Messages }}

{{/* Variables */}}
{{- $fieldNameFormat := $message.GoMax | printf "%%-%ds " -}}

type {{ $message.Name }} struct {
{{- range $field := $message.Fields }}
	{{- /* Field name */}}
	{{ $field.GoName | printf $fieldNameFormat }}
	{{- /* Field type */}}
	{{- if $field.GoIsArray }}[]{{ end }}
	{{- if $field.GoIsPointer }}*{{ end }}
	{{- $field.GoType }}
	{{- /* Field tags */}}
	{{- if len $field.GoTags }} ` + "`" + `{{ end }}
	{{- range $index, $tag := $field.GoTags }}
		{{- $tag.Name }}:"{{ $tag.Value }}"
		{{- if not (last $index $field.GoTags) }} {{ end }}
	{{- end }}
	{{- if len $field.GoTags }}` + "`" + `{{ end }}
{{- end }}
}
{{- end }}
`