{{ /* Trigger: Command: lore */ }}

{{$endpoint:= "https://www.ishtar-collective.net/search/"}}
{{$suffix:= ""}}
{{$searchurl:= ""}}
{{$query:= ""}}
{{$argslength:= len .CmdArgs}}
{{$args := (joinStr " " .CmdArgs)}}
{{if gt $argslength 1}}
	{{$firstarg:= (index .CmdArgs 0)}}
	{{if eq $firstarg "categories"}}
		{{$args = (joinStr " " (slice .CmdArgs 1))}}
		{{$suffix = "/page/1?document_type_id=0"}}
	{{else if eq $firstarg "books"}}
		{{$args = (joinStr " " (slice .CmdArgs 1))}}
		{{$suffix = "/page/1?document_type_id=0"}}
	{{else if eq $firstarg "items"}}
		{{$args = (joinStr " " (slice .CmdArgs 1))}}
		{{$suffix = "/page/1?document_type_id=6"}}
	{{end}}
{{end}}
{{if eq $argslength 0}}
Please provide a valid query.
{{else}}
	{{$query = urlescape $args}}
	{{$searchurl = (joinStr "" $endpoint "" $query "" $suffix "")}}
	{{$searchurl}}
{{end}}