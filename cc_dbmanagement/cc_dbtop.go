{{ /* Trigger: Exact Match: ~dbtop */ }}
{{ /* Reports top 100 databse entries for all keys and users. */ }}

{{$lb := dbTopEntries "%" 100 0}}
{{range $lb}}
`{{.UserID}}` **:** `{{.Key}}` **:** `{{.Value}}`
{{end}}
done!