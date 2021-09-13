{{ $balanceKey := "RSCoinBalance" }}
{{ $gameKey := "RSCasinoDuel" }}
{{ $e := "<:RSStonkCoin:869340420692394095>" }}
{{ $historyKey := "RSCasinoDuelHistory" }}

{{ $x := .Reaction.MessageID }}
{{ $gameState := (dbGet $x $gameKey).Value }}
{{ $user := .User }}

{{if eq (printf "%T" $gameState) "*templates.SDict"}}
	{{ $state := (toInt ($gameState.Get "state"))}}
	{{ $r := .Reaction.Emoji.Name }}
	{{ $amount := $gameState.Get "bet" }}

	{{ $user1 := userArg ($gameState.Get "user1")}}
	{{ $user2 := userArg ($gameState.Get "user2")}}
	{{ $user1balance := toInt (dbGet $user1.ID $balanceKey).Value }}
	{{ $user2balance := toInt (dbGet $user2.ID $balanceKey).Value }}

	{{ if eq 0 $state }}
		{{/* Waiting for game to start */}}

		{{ if eq $user.ID $user2.ID }}
			{{ if eq $r "✅"}}
				{{ deleteAllMessageReactions nil $x }}
				{{ $gameState.Set "health1" 100 }}
				{{ $gameState.Set "health2" 100 }}
				{{ $gameState.Set "state" 1 }}
				{{ dbSetExpire $x $gameKey $gameState 60}}

				{{ $hlth1 := ($gameState.Get "health1")}}
				{{ $hlth2 := ($gameState.Get "health2")}}

				{{$embed := cembed
					"title" (joinStr "" "__" $user1.Username "__ challenges __" $user2.Username "__ to a duel!")
					"description" "The game is on!"
					"fields" (cslice
					(sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true )
					(sdict "name" $user1.Username "value" (joinStr "" $hlth1 "♥" ) "inline" true )
					(sdict "name" $user2.Username "value" (joinStr "" $hlth2 "♥" ) "inline" true ))
				}}
				{{ editMessage nil $x (complexMessageEdit "content" $user1.Mention "embed" $embed) }}
				{{ addMessageReactions nil $x "⚔" "🛡" }}
			{{ else if eq $r "❌" }}
				{{ deleteAllMessageReactions nil $x }}
				{{$embed := cembed
					"title" (joinStr "" "__" $user1.Username "__ challenged __" $user2.Username "__ to a duel.")
					"description" "The duel was declined."
				}}
				{{ dbDel $user1.ID $gameKey }}
				{{ dbDel $user2.ID $gameKey }}
				{{ dbDel $x $gameState }}
				{{ editMessage nil $x (complexMessageEdit "content" " " "embed" $embed) }}
			{{ end }}
		{{ end }}

	{{ else }}
		{{/*Game has started */}}
		{{ $curUser := $user1 }}
		{{ $opponent := $user2 }}
		{{ $dmgStr := "dmg1" }}
		{{ if or (eq $state 2) (eq $state 4)}} {{ $curUser = $user2 }} {{ $dmgStr = "dmg2"}} {{ $opponent = $user1 }} {{ end }}
		{{/*CurUser is defined as person whose turn it is*/}}

		{{ if eq $user.ID $curUser.ID }}
			{{/*Current user is correct*/}}
			{{ $actionDesc := " " }}
			
			{{ $hlth1 := ($gameState.Get "health1")}}
			{{ $hlth2 := ($gameState.Get "health2")}}

			{{ $dmg := (randInt 15 | add 15)}}
			{{ if eq $curUser.ID $user.ID}}
				{{/*Reacting user is correct*/}}
				{{ if eq $r "⚔"}}
					{{ $actionDesc = "attacks!"}}
					{{ $gameState.Set $dmgStr $dmg }}
				{{ else if eq $r "🛡"}}
					{{ $actionDesc = "braces!"}}
					{{ $gameState.Set $dmgStr (mult $dmg -1) }}
				{{ end }}
			{{ end }}
			
			{{ $embedFields := (cslice
				(sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true ))
			}}

			{{ $state = add $state 1 }}
			{{ if or (eq $state 3) (eq $state 6)}}
				{{ $state = add $state 1 }}
				{{/*Game is at calculation state*/}}
				{{ deleteAllMessageReactions nil $x }}

				{{ $dmg1 := toInt ($gameState.Get "dmg1")}}
				{{ $dmg2 := toInt ($gameState.Get "dmg2")}}
				{{ $parry := (add $dmg2 $dmg1) }}

				{{$aStr1 := "*Blocked!*"}}
				{{$aStr2 := "*Blocked!*"}}
				
				{{ if gt $dmg1 0 }}
					{{ $aStr1 = (joinStr "" "Attacks for `" $dmg1 "`") }}
					{{ if lt $dmg2 0 }}
						{{/*Target is blocking*/}}
						{{ if le $parry 0 }}
							{{ $parry = (mult $parry 3)}}
							{{ if eq $parry 0 }} {{ $parry = -1 }} {{ end }}
							{{ $hlth1 = add $hlth1 $parry }}
							{{ $aStr2 = (joinStr "" "Parried for `" (mult $parry -1) "`")}}
							{{ $aStr1 = "Got parried!"}}
						{{ else }}
							{{ $aStr2 = "Failed to block!"}}
							{{ $hlth2 = sub $hlth2 $dmg1 }}
						{{ end }}
					{{ else }}
						{{ $hlth2 = sub $hlth2 $dmg1 }}
					{{ end }}
				{{ end }}

				{{ if gt $dmg2 0 }}
					{{ $aStr2 = (joinStr "" "Attacks for `" $dmg2 "`") }}
					{{ if lt $dmg1 0 }}
						{{/*Target is blocking*/}}
						{{ if le $parry 0 }}
							{{ $parry = (mult $parry 3)}}
							{{ if eq $parry 0 }} {{ $parry = -1 }} {{ end }}
							{{ $hlth2 = add $hlth2 $parry }}
							{{ $aStr1 = (joinStr "" "Parried for `" (mult $parry -1) "`")}}
							{{ $aStr2 = "Got parried!"}}
						{{ else }}
							{{ $aStr1 = "Failed to block!"}}
							{{ $hlth1 = sub $hlth1 $dmg2 }}
						{{ end }}
					{{ else }}
						{{ $hlth1 = sub $hlth1 $dmg2 }}
					{{ end }}
				{{ end }}

				{{ if and (lt $dmg1 0) (lt $dmg2 0) }}
					{{ $dmg = (randInt 12 | add 3 )}}
					{{ $aStr1 = (joinStr "" "Took `" $dmg "`")}}
					{{ $aStr2 = (joinStr "" "Took `" $dmg "`")}}
					{{ $hlth1 = sub $hlth1 $dmg }}
					{{ $hlth2 = sub $hlth2 $dmg }}
				{{ end }}

				{{ $embedFields = $embedFields.AppendSlice (cslice (sdict "name" $user1.Username "value" $aStr1 "inline" true ) 
					(sdict "name" $user2.Username "value" $aStr2 "inline" true ) 
					(sdict "name" "Results:" "value" "`After last round:`" "inline" true )
					(sdict "name" $user1.Username "value" (joinStr "" $hlth1 "♥" ) "inline" true ) 
					(sdict "name" $user2.Username "value" (joinStr "" $hlth2 "♥" ) "inline" true )
					(sdict "name" (joinStr "" $curUser.Username " " $actionDesc ) "value" "is back up!" "inline" false ))
				}}
				
				{{ $gameState.Set "health1" $hlth1 }}
				{{ $gameState.Set "health2" $hlth2 }}

				{{ if ge $state 6 }} {{ $state = 1 }} {{ end }}

				{{ if or (le $hlth1 0) (le $hlth2 0) }}
					{{ $state = 7 }}
					{{ $gameState.Set "state" $state }}
				{{ else }}
					{{ addMessageReactions nil $x "⚔" "🛡" }}
				{{ end }}
			{{ else }}
				{{ deleteAllMessageReactions nil $x }}
				{{ addMessageReactions nil $x "⚔" "🛡" }}
				{{ $embedFields = $embedFields.AppendSlice (cslice 
					(sdict "name" $user1.Username "value" (joinStr "" $hlth1 "♥" ) "inline" true ) 
					(sdict "name" $user2.Username "value" (joinStr "" $hlth2 "♥" ) "inline" true )
					(sdict "name" (joinStr "" $curUser.Username " " $actionDesc ) "value" (joinStr "" $opponent.Username " is up!") "inline" false ))
				}}
			{{ end }}

			{{ $gameState.Set "state" $state }}

			{{$actionEmbed := cembed
				"title" (joinStr "" "__" $user1.Username "__ challenges __" $user2.Username "__ to a duel!")
				"description" "The game is on!"
				"fields" $embedFields 
			}}
			{{ if or (eq ($gameState.Get "state") 2) (eq ($gameState.Get "state") 4)}} 
				{{ editMessage nil $x (complexMessageEdit "content" $user2.Mention "embed" $actionEmbed) }}
			{{ else }}
				{{ editMessage nil $x (complexMessageEdit "content" $user1.Mention "embed" $actionEmbed) }}
			{{ end }}
			
			{{ dbSetExpire $x $gameKey $gameState 60 }}
			{{ dbSetExpire $user1.ID $gameKey (sdict "active" 1 "messageID" $x) 60 }}
			{{ dbSetExpire $user2.ID $gameKey (sdict "active" 1 "messageID" $x) 60 }}

			{{ if eq $state 7}}
				{{/*Game is over, shut it down and hand out winnings.*/}}
				{{ deleteAllMessageReactions nil $x }}
				{{ sleep 5 }}
				{{ $winFields := cslice (sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true ) }}

				{{ $user1History := (dbGet $user1.ID $historyKey).Value }}
				{{ $user2History := (dbGet $user2.ID $historyKey).Value }}
				{{if not (eq (printf "%T" $user1History) "*templates.SDict")}} {{ $user1History = (sdict "wins" 0 "losses" 0 )}} {{ end }}
				{{if not (eq (printf "%T" $user2History) "*templates.SDict")}} {{ $user2History = (sdict "wins" 0 "losses" 0 )}} {{ end }}

				{{ if and (le $hlth1 0) (le $hlth2 0)}}
					{{ $winFields = $winFields.Append (sdict "name" "Game Over!" "value" "It was a draw!"  "inline" true )}}
				{{ else if gt $hlth1 $hlth2}}
					{{/*User 1 wins*/}}
					{{ $winFields = $winFields.Append (sdict "name" "Game Over!" "value" (joinStr "" $user1.Username " won!")  "inline" true )}}
					{{ $user1balance = (add $user1balance $amount )}}
					{{ $user2balance = (sub $user2balance $amount )}}
					
					{{ $user1History.Set "wins" (add (toInt ($user1History.Get "wins")) 1)}}
					{{ $user2History.Set "losses" (add (toInt ($user2History.Get "losses")) 1)}}
				{{ else }}
					{{/*User 2 wins*/}}
					{{ $winFields = $winFields.Append (sdict "name" "Game Over!" "value" (joinStr "" $user2.Username " won!")  "inline" true )}}
					{{ $user1balance = (sub $user1balance $amount )}}
					{{ $user2balance = (add $user2balance $amount )}}

					{{ $user1History.Set "losses" (add (toInt ($user1History.Get "losses")) 1)}}
					{{ $user2History.Set "wins" (add (toInt ($user2History.Get "wins")) 1)}}
				{{ end }}

				{{ dbSet $user1.ID $balanceKey $user1balance }}
				{{ dbSet $user2.ID $balanceKey $user2balance }}
				
				{{ dbSet $user1.ID $historyKey $user1History}}
				{{ dbSet $user2.ID $historyKey $user2History}}

				{{$embed := cembed
					"title" (joinStr "" "__" $user1.Username "__ challenged __" $user2.Username "__ to a duel!")
					"fields" $winFields
					"footer" (sdict "text" (joinStr "" $user1.Username " has " $user1balance " coins | " $user2.Username " has " $user2balance " coins") )
				}}
				{{ editMessage nil $x (complexMessageEdit "content" " " "embed" $embed) }}
				
				{{ $silent := dbDel $user1.ID $gameKey }}
				{{ $silent = dbDel $user2.ID $gameKey }}
				{{ $silent = dbDel $x $gameKey }}
				
			{{ end }}
		{{ end }}
	{{ end }}
{{ end }}