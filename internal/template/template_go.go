package template

const defaultGoTemplate = `/* Generated by protoc-gen-tpl. Do not edit! */

package {{ .GoPackage }}
{{- if len .GoImports }}

import (
{{- range $import := .GoImports }}
	{{ $import }}
{{- end }}
)

{{- end }}
{{- range $message := .Messages }}

{{/* Variables */}}
{{- $fieldNameFormat := $message.GoMaxName | printf "%%-%ds " -}}
{{- $fieldTypeFormat := $message.GoMaxType | printf "%%-%ds" -}}

type {{ $message.Name }} struct {
{{- range $field := $message.Fields }}
	{{- /* Field name */}}
	{{ $field.GoName | printf $fieldNameFormat }}
	{{- /* Field type */}}
	{{- $field.GoType | printf $fieldTypeFormat }}
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
