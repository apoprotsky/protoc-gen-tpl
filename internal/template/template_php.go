package template

const defaultPhpTemplate = `<?php
/* Generated by protoc-gen-tpl. Do not edit! */

namespace {{ .PhpPackage }};

class {{ .Message.Name }}
{
{{- range $field := .Message.Fields }}
    public {{/**/}}
    {{- /* Field type */ -}}
    {{ $field.PhpType }} {{/**/}}
    {{- /* Field name */ -}}
    ${{ $field.PhpName }};
{{- end }}
}
`
