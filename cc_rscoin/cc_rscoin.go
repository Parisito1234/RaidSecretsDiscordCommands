{{$perms := "ManageServer"}}
{{ $key := "RSCoinBalance" }}
{{ $claimKey := "RSCoinStarterClaimed" }}
{{ $coinIcon := "<:RSStonkCoin:869340420692394095>" }}
{{ $args := parseArgs 1 "Syntax is `<action> <user> <amount>` - only users with `Manage Server` can modify other user balances."
	(carg "string" "action - balance, set, add, remove")
	(carg "userid" "user to perform action on")
	(carg "int" "amount") }}
{{ $action := lower ($args.Get 0) }}
{{ $lotteryPool := toInt ((dbGet 204255221017214977 $key).Value) }}
 

{{/* If user is not defined in arg 1, set command's user to targetUser*/}}
{{ $targetUser := $.User }}
{{ $curBalance := toInt (dbGet $targetUser.ID $key).Value }}


{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	{{/* User has manage server perms! */}}

	{{ if $args.IsSet 1}} 
		{{/* User is defined in arg 1, save to targetUser*/}}
		{{ $targetUser = (userArg ($args.Get 1) ) }}
	{{ end }}
	{{ $curBalance := toInt (dbGet $targetUser.ID $key).Value }}

	{{ if or (eq $action "balance") (eq $action "bal") }}
		{{/* Balance, with perms! */}}
		{{ $curBalance := toInt (dbGet $targetUser.ID $key).Value }}
		{{ sendMessage nil (joinStr "" "`" $targetUser.Username "` has `" $curBalance "` " $coinIcon )}}
	{{ else if or (eq $action "add") (eq $action "give") }} 
		{{ if not ($args.Get 2) }}
			{{ sendMessage nil "Missing or incorrect format for amount. `-rscoin add|give @user amount`" }}
		{{ else }}
			{{ $value := toInt ($args.Get 2) }}
			{{ $newBalance := add $curBalance $value }}
			{{ dbSet $targetUser.ID $key (str $newBalance) }}
			{{ sendMessage nil (joinStr "" "`" $targetUser.Username "` has `" (toInt (dbGet $targetUser.ID $key).Value ) "` " $coinIcon )}}
		{{ end }}
	{{ else if or (eq $action "remove") (eq $action "take") }}
		{{ if not ($args.Get 2) }}
			{{ sendMessage nil "Missing or incorrect format for amount. `-rscoin remove|take @user amount`" }}
		{{ else }}
			{{ $value := toInt ($args.Get 2) }}
			{{ $newBalance := sub $curBalance $value }}
			{{ if lt $newBalance 0 }} {{$newBalance = 0 }} {{ end }}
			{{ dbSet $targetUser.ID $key (str $newBalance) }}
			{{ sendMessage nil (joinStr "" "`" $targetUser.Username "` has `" (toInt (dbGet $targetUser.ID $key).Value ) "` " $coinIcon )}}
		{{ end }}
	{{ else if eq $action "set" }}
		{{/* Set balance! */}}
		{{ if not ($args.Get 2) }}
			{{ sendMessage nil "Missing or incorrect format for amount. `-rscoin set @user amount`" }}
		{{ else }}
			{{ $value := ($args.Get 2) }}
			{{ if lt $value 0 }}
				{{ sendMessage nil "Can't set a value lower than 0." }}
			{{ else }}
				{{ dbSet $targetUser.ID $key  (str $value) }}
				{{ sendMessage nil (joinStr "" "`" $targetUser.Username "` has `" (toInt (dbGet $targetUser.ID $key).Value ) "` " $coinIcon )}}
			{{ end }}
		{{ end }}
	{{ else if or (eq $action "pool") (eq $action "lotto") (eq $action "lottery") }}
		{{ sendMessage nil (joinStr "" "Current Lottery Pool balance is `" $lotteryPool "` " $coinIcon ) }}
	{{ else if eq $action "resetlotto" }}
		{{ dbSet 204255221017214977 $key 0 }}
		{{ sendMessage nil (joinStr "" "Current Lottery Pool balance is `0` " $coinIcon ) }}
	{{ else if or (eq $action "top") (eq $action "list") }}
		{{ $userList := cslice }}
		{{ $dbtop := dbTopEntries $key 20 0 }}
		{{ range $dbtop }}
			{{ $currentMember := (getMember .UserID) }}
			{{ $roles := ($.Guild.GetMemberPermissions $.Channel.ID $currentMember.User.ID $currentMember.Roles) }}
			{{ if and (ge $roles 0) (lt $roles 2140775575) }}
				{{ $userList = $userList.Append $currentMember.User.ID }}
			{{ end }}
		{{ end }}
		{{ $displayList := cslice "" "" "" "" "" }}
		{{ range $i, $element := (seq 0 5) }}
			{{ $displayList.Set $i (index $userList $i) }}
		{{ end }}

		{{$embed := cembed 
			"title" (joinStr "" "__" $coinIcon " Leaderboards__")
			"color" 1772743
			"description" (joinStr "" "**1:** `" (userArg (index $displayList 0)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 0) $key).Value) "`"
			"\n**2:** `" (userArg (index $displayList 1)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 1) $key).Value) "`"
			"\n**3:** `" (userArg (index $displayList 2)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 2) $key).Value) "`"
			"\n**4:** `" (userArg (index $displayList 3)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 3) $key).Value) "`"
			"\n**5:** `" (userArg (index $displayList 4)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 4) $key).Value) "`" )
			"footer" (sdict "text" (joinStr "" "Current Lottery Balance: " $lotteryPool ))
		}}
		{{ sendMessage nil $embed }}
		{{ deleteTrigger 0 }}
	{{ end }}


{{ else }}
	{{/* User does not have manage server perms! */}}

	{{ $curBalance := toInt (dbGet $.User.ID $key).Value }}
	{{ if or (eq $action "balance") (eq $action "bal") }}
		{{ sendMessage nil (joinStr "" "`" $.User.Username "` has `" $curBalance "` " $coinIcon )}}
	{{ else if or (eq $action "top") (eq $action "list") }}
		{{ $userList := cslice }}
		{{ $dbtop := dbTopEntries $key 20 0 }}
		{{ range $dbtop }}
			{{ $currentMember := (getMember .UserID) }}
			{{ $roles := ($.Guild.GetMemberPermissions $.Channel.ID $currentMember.User.ID $currentMember.Roles) }}
			{{ if and (ge $roles 0) (lt $roles 2140775575) }}
				{{ $userList = $userList.Append $currentMember.User.ID }}
			{{ end }}
		{{ end }}
		{{ $displayList := cslice "" "" "" "" "" }}
		{{ range $i, $element := (seq 0 5) }}
			{{ $displayList.Set $i (index $userList $i) }}
		{{ end }}

		{{$embed := cembed 
			"title" (joinStr "" "__" $coinIcon " Leaderboards__")
			"color" 1772743
			"description" (joinStr "" "**1:** `" (userArg (index $displayList 0)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 0) $key).Value) "`"
			"\n**2:** `" (userArg (index $displayList 1)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 1) $key).Value) "`"
			"\n**3:** `" (userArg (index $displayList 2)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 2) $key).Value) "`"
			"\n**4:** `" (userArg (index $displayList 3)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 3) $key).Value) "`"
			"\n**5:** `" (userArg (index $displayList 4)).Username "` : " $coinIcon " `" (toInt (dbGet (index $displayList 4) $key).Value) "`" )
			"footer" (sdict "text" (joinStr "" "Current Lottery Balance: " $lotteryPool ))
		}}
		{{ sendMessage nil $embed }}
		{{ deleteTrigger 0 }}
	{{ else if or (eq $action "give") (eq $action "pay")}}
		{{ if not ($args.Get 2) }}
			{{ sendMessage nil "Missing or incorrect format for amount. `-rscoin <give|pay> @user amount`" }}
		{{ else }}
			{{ $amount := toInt ($args.Get 2) }}
			{{ if gt $amount $curBalance}}
				{{ $amount = $curBalance}}
			{{ end }}
			{{ if lt $amount 0 }} {{ $amount = mult $amount -1 }} {{ end }}

			{{ $givingUser := $.User }}
			{{ $givingBalance := (toInt (dbGet $givingUser.ID $key).Value) }}
			{{ $newGivingBalance := (sub $givingBalance $amount)}}
			{{ dbSet $givingUser.ID $key $newGivingBalance }}

			{{ $targetUser = (userArg ($args.Get 1) ) }}
			{{ $targetBalance := (toInt (dbGet $targetUser.ID $key).Value) }}
			{{ $newTargetBalance := (add $targetBalance $amount)}}
			{{ dbSet $targetUser.ID $key $newTargetBalance }}

			{{ sendMessage nil (joinStr "" "`" $givingUser.Username "` paid `" $targetUser.Username "` `" $amount "`" $coinIcon) }}
			{{ sendMessage nil (joinStr "" $targetUser.Username " " $targetBalance " " (dbGet $targetUser.ID $key).Value " " $givingUser.Username " " $givingBalance " " (dbGet $givingUser.ID $key).Value " " $amount )}}

		{{ end }}
	{{ else if eq $action "claim" }}
		{{ $claimState := toInt (dbGet $.User.ID $claimKey).Value }}
		{{ if eq $claimState 0 }}
			{{ dbSet $.User.ID $claimKey 1}}				
			{{ dbSet $.User.ID $key (str (add $curBalance 10)) }}
			{{ sendMessage nil (joinStr "" "`" $.User.Username "` has `" (add $curBalance 10) "` " $coinIcon )}}
		{{ else }}
			{{ sendMessage nil "You have already claimed your free RSCoin" }}
		{{end}}
	{{ else if or (eq $action "empty") (eq $action "dump") (eq $action "remove") }}
		{{ $value := $curBalance }}
		{{ if not ($args.Get 1) }}
			{{ dbDel $.User.ID $key }}
		{{ else }}
			{{ $value = toInt ($args.Get 1) }}
			{{ if lt $value 0 }} {{ $value = mult $value -1 }} {{ end }}
			{{ if gt $value $curBalance}}
				{{ $value = $curBalance}}
			{{ end }}
			{{ $newBalance := sub $curBalance $value }}
			{{ if lt $newBalance 0 }} {{$newBalance = 0 }} {{ end }}
			{{ dbSet $.User.ID $key (str $newBalance) }}
		{{ end }}
		{{ sendMessage nil (joinStr "" "`" $.User.Username "` just dumped " $value " " $coinIcon " into the pool!") }}
		{{ dbSet 204255221017214977 $key (add $lotteryPool $value) }}
	{{ else if or (eq $action "pool") (eq $action "lotto") (eq $action "lottery")}}
		{{ sendMessage nil (joinStr "" "Current Lottery Pool balance is `" $lotteryPool "` " $coinIcon ) }}
	{{ else }}
		Improper syntax or missing permissions.
	{{ end }}
{{end}}
