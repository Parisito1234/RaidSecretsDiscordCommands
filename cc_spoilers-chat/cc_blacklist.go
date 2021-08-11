{{ /* Trigger: Command: blacklist */ }}
{{ /* Channel/User Restrictions: Require role: Moderator */ }}
{{ /* Removes user from being able to talk inside of the spoilers channel, works with a RoleMenu set to ignore "Blacklist" role when reacting for the access role */ }}

{{ $argslength := len .CmdArgs }}
{{ $whitelistrole := CHANNEL_WHITELIST_ROLE_ID }}
{{ $blacklistrole := CHANNEL_BLACKLIST_ROLE_ID }}
{{ $modlogchannel := MODLOG_CHANNEL_ID }}


{{ if eq $argslength 0 }}
    {{ sendMessage $modlogchannel "Blacklist command Usage: `-blacklist <UserID | @mention>`" }}
{{ else }}
    {{ $user := (userArg (index .CmdArgs 0)) }}
    {{ $author := .Message.Author }}

    {{ if targetHasRoleID $user $blacklistrole}}
        {{sendMessage $modlogchannel "Target user is already blacklisted in spoiler chat." }}
    {{ else }}
        {{ if targetHasRoleID $user $whitelistrole }}
            {{ takeRoleID $user.ID $whitelistrole }}
        {{ end}}
        
        {{ giveRoleID $user.ID $blacklistrole }}

        {{ $userString := (joinStr "" ":x: Blacklisted " $user.String " from #spoilers." ) }}
        {{ $userInfo := (joinStr "" "*(ID" $user.ID ")*" ) }}
        {{ $authorString := (joinStr "" $author.String " (ID " $author.ID ")" ) }}
        {{$embed := cembed 
            "color" 13852449
            "thumbnail" (sdict "url" ($user.AvatarURL "256") )
            "author" (sdict "name" $authorString "icon_url" ($author.AvatarURL "64") )
            "fields" (cslice
                (sdict "name" $userString "value" $userInfo )
            )
        }}

        {{ sendMessage $modlogchannel $embed }}
    {{ end }}
{{ end }}
{{ deleteTrigger 0 }}