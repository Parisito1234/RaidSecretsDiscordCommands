{{ /* Trigger: Exact Match: ~dbbottom */ }}
{{ /* Reports bottom 100 databse entries for all keys and users. */ }}


{{$lb := dbBottomEntries "%" 2 0}}
{{range $lb}}
`{{.UserID}}` **:** `{{.Key}}` **:** `{{.Value}}`
{{end}}
done!