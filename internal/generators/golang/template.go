package golang

const defaultTemplate = `/* Generated by protoc-gen-tpl. Do not edit! */

package {{ .Package }}
{{- range $struct := .Structs }}

{{/* Variables */}}
{{- $fieldNameFormat := $struct.MaxFieldNameLength | printf "%%-%ds " -}}

type {{ $struct.Name }} struct {
{{- range $field := $struct.Fields }}
	{{- /* Field name */}}
	{{ $field.Name | printf $fieldNameFormat }}
	{{- /* Field type */}}
	{{- if $field.IsArray }}[]{{ end }}
	{{- if $field.IsPointer }}*{{ end }}
	{{- $field.Type }}
	{{- /* Field tags */}}
	{{- if len $field.Tags }} ` + "`" + `{{ end }}
	{{- range $index, $tag := $field.Tags }}
		{{- $tag.Name }}:"{{ $tag.Value }}"
		{{- if not (last $index $field.Tags) }} {{ end }}
	{{- end }}
	{{- if len $field.Tags }}` + "`" + `{{ end }}
{{- end }}
}
{{- end }}
`
