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
		{{ sendMessage nil (joinStr "" "`" $targetUser "` has " $curBalance " " $coinIcon )}}
	{{ else if or (eq $action "add") (eq $action "give") }} 
		{{ if not ($args.Get 2) }}
			{{ sendMessage nil "Missing or incorrect format for amount. `-rscoin add|give @user amount`" }}
		{{ else }}
			{{ $value := toInt ($args.Get 2) }}
			{{ $newBalance := add $curBalance $value }}
			{{ dbSet $targetUser.ID $key (str $newBalance) }}
			{{ sendMessage nil (joinStr "" "`" $targetUser "` has " (toInt (dbGet $targetUser.ID $key).Value ) " " $coinIcon )}}
		{{ end }}
	{{ else if or (eq $action "remove") (eq $action "take") }}
		{{ if not ($args.Get 2) }}
			{{ sendMessage nil "Missing or incorrect format for amount. `-rscoin remove|take @user amount`" }}
		{{ else }}
			{{ $value := toInt ($args.Get 2) }}
			{{ $newBalance := sub $curBalance $value }}
			{{ if lt $newBalance 0 }} {{$newBalance = 0 }} {{ end }}
			{{ dbSet $targetUser.ID $key (str $newBalance) }}
			{{ sendMessage nil (joinStr "" "`" $targetUser "` has " (toInt (dbGet $targetUser.ID $key).Value ) " " $coinIcon )}}
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
				{{ sendMessage nil (joinStr "" "`" $targetUser "` has " (toInt (dbGet $targetUser.ID $key).Value ) " " $coinIcon )}}
			{{ end }}
		{{ end }}
	{{ else if or (eq $action "pool") (eq $action "lotto") (eq $action "lottery") }}
		{{ sendMessage nil (joinStr "" "Current Lottery Pool balance is `" $lotteryPool "` " $coinIcon ) }}
	{{ else if eq $action "resetlotto" }}
		{{ dbSet 204255221017214977 $key 0 }}
		{{ sendMessage nil (joinStr "" "Current Lottery Pool balance is `0` " $coinIcon ) }}
	{{ else if or (eq $action "top") (eq $action "list") }}
		{{ $dbtop := dbTopEntries $key 5 0 }}
		{{ range $dbtop }}
			`{{(userArg .UserID).String}}` - `{{.Key}}` : `{{.Value}}`
		{{ end }}
	{{ end }}


{{ else }}
	{{/* User does not have manage server perms! */}}

	{{ $curBalance := toInt (dbGet $.User.ID $key).Value }}
	{{ if or (eq $action "balance") (eq $action "bal") }}
		{{ sendMessage nil (joinStr "" "`" $.User "` has " $curBalance " " $coinIcon )}}
	{{ else if or (eq $action "top") (eq $action "list") }}
		{{ $dbtop := dbTopEntries $key 5 0 }}
		{{ range $dbtop }}
			`{{(userArg .UserID).String}}` - `{{.Key}}` : `{{printf "%v" .Value}}`
		{{ end }}
	{{ else if or (eq $action "give") (eq $action "pay")}}
		{{ if not ($args.Get 2) }}
			{{ sendMessage nil "Missing or incorrect format for amount. `-rscoin give|pay @user amount`" }}
		{{ else }}
			{{/*Take calling user's balance (user) and move $value to $targetUser's balance*/}}
			
			{{ $targetUser = (userArg ($args.Get 1) ) }}
			{{ $targetBalance := toInt (dbGet $targetUser.ID $key).Value }}
			{{ $callingUser := $.User }}
			{{ $callingBalance := toInt (dbGet $callingUser.ID $key).Value }}
			{{ $value := toInt ($args.Get 2) }}
			{{ if gt $value $curBalance}}
				{{ $value = $curBalance}}
			{{ end }}
			{{ if lt $value 0 }} {{ $value = mult $value -1 }} {{ end }}
			{{ dbSet $targetUser.ID $key (add $targetBalance $value) }}
			{{ dbSet $callingUser.ID $key (sub $callingBalance $value) }}
			{{ sendMessage nil (joinStr "" $callingUser.String " paid " $targetUser.String " " $value $coinIcon) }}

		{{ end }}
	{{ else if eq $action "claim" }}
		{{ $claimState := toInt (dbGet $.User.ID $claimKey).Value }}
		{{ if eq $claimState 0 }}
			{{ dbSet $.User.ID $claimKey 1}}				
			{{ dbSet $.User.ID $key (str (add $curBalance 10)) }}
			{{ sendMessage nil (joinStr "" "`" $.User "` has " (add $curBalance 10) " " $coinIcon )}}
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
		{{ sendMessage nil (joinStr "" "`" $.User "` just dumped " $value " " $coinIcon " into the pool!") }}
		{{ dbSet 204255221017214977 $key (add $lotteryPool $value) }}
	{{ else if or (eq $action "pool") (eq $action "lotto") (eq $action "lottery")}}
		{{ sendMessage nil (joinStr "" "Current Lottery Pool balance is `" $lotteryPool "` " $coinIcon ) }}
	{{ else }}
		Improper syntax or missing permissions.
	{{ end }}
{{end}}
