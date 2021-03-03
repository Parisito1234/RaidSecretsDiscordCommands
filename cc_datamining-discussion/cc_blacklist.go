{{ $argslength := len .CmdArgs }}
{{ $whitelistrole := 813768536420909126 }}
{{ $blacklistrole := 813771735865360384 }}
{{ if eq $argslength 0 }}
	{{ sendMessage nil "Usage: `-blacklist <UserID | @mention>`" }}
{{ else }}
	{{ $user := (userArg (index .CmdArgs 0)) }}
	{{ if targetHasRoleID $user $blacklistrole}}
		{{sendMessage nil "Target user is already blacklisted." }}
	{{ else }}
		{{ if targetHasRoleID $user $whitelistrole }}
			{{ takeRoleID $user.ID $whitelistrole }}
		{{ end}}
		
		{{ giveRoleID $user.ID $blacklistrole }}
		{{ sendMessage nil (joinStr "" "User: " "`" $user.Username "`is no longer able to talk in this channel" "") }}
	{{ end }}
{{ end }}