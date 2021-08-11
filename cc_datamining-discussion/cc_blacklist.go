{{ /* Trigger: Command: blacklist */ }}
{{ /* Channel/User Restrictions: Require role: Moderator or Dataminer */ }}
{{ /* Removes user from being able to talk inside of the datamining channel, works with a RoleMenu set to ignore "Blacklist" role when reacting for the access role*/ }}

{{ $argslength := len .CmdArgs }}
{{ $whitelistrole := CHANNEL_WHITELIST_ROLE_ID }}
{{ $blacklistrole := CHANNEL_BLACKLIST_ROLE_ID }}
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