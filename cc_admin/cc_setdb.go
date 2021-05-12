{{ $args := (joinStr " " .CmdArgs) }}
{{ $argslength := len .CmdArgs }}

{{ if gt $argslength 1 }}
	{{ $dbkey := (index .CmdArgs 0) }}
	{{ $dbvalue := (joinStr " " (slice .CmdArgs 1)) }}
	{{ dbSet 0 $dbkey $dbvalue }}
	{{ sendMessage nil (dbGet 0 $dbkey).Value }}
{{ else }}
	{{ sendMessage nil "Usage: `-setdb <key> <value>`" }}
{{ end }}
