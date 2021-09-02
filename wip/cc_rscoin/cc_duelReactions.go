{{ $balanceKey := "RSCoinBalance" }}
{{ $gameKey := "RSCasinoDuel" }}
{{ $e := "<:RSStonkCoin:869340420692394095>" }}
{{ $historyKey := "RSCasinoDuelHistory" }}

{{ $user := .User }}
{{ $x := .Reaction.MessageID }}
{{ $gameState := (dbGet $x $gameKey).Value }}
{{ $state := (toInt ($gameState.Get "state"))}}
{{ $r := .Reaction.Emoji.Name }}
{{ $user1 := userArg ($gameState.Get "user1")}}
{{ $user2 := userArg ($gameState.Get "user2")}}
{{ $amount := $gameState.Get "bet" }}


{{if eq (printf "%T" $gameState) "*templates.SDict"}}
	{{ if eq 0 $state }}
		{{/* Waiting for game to start */}}

		{{ if eq $user.ID $user2.ID }}
			{{ if eq $r "✅"}}
				{{ deleteAllMessageReactions nil $x }}
				{{ $gameState.Set "health1" 100 }}
				{{ $gameState.Set "health2" 100 }}
				{{ $gameState.Set "state" 1 }}
				{{ dbSetExpire $x $gameKey $gameState 60}}

				{{ $health1 := ($gameState.Get "health1")}}
				{{ $health2 := ($gameState.Get "health2")}}

				{{$embed := cembed
					"title" (joinStr "" "__" $user1.Username "__ challenges __" $user2.Username "__ to a duel!")
					"description" "The game is on!"
					"fields" (cslice
					(sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true )
					(sdict "name" $user1.Username "value" (joinStr "" $health1 "♥" ) "inline" true )
					(sdict "name" $user2.Username "value" (joinStr "" $health2 "♥" ) "inline" true ))
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
			
		{{ else if eq $user.ID $user1.ID }}
			{{ deleteMessageReaction nil $x .Reaction.UserID $r }}
			{{ sendMessage nil "Wait for your opponent to start the match!"}}
		{{ end }}

	{{ else if ge $state 1 }}
		{{/*Game has started */}}
		{{ $curUser := $user1 }}
		{{ $opponent := $user2 }}
		{{ if eq ($gameState.Get "state") 2}} {{ $curUser = $user2 }} {{ $opponent = $user1 }} {{ end }}
		{{/*CurUser is defined as person whose turn it is*/}}

		{{ if eq $user.ID $curUser.ID }}
			{{/*Current user is correct*/}}
			
			{{ $actionDesc := " " }}
			{{ $dmg := (randInt 15 | add 15)}}
			{{ if eq $curUser.ID $user.ID}}
				{{/*Reacting user is correct*/}}
				{{ if eq $r "⚔"}}
					{{ $actionDesc = "attacks!"}}
					{{ $gameState.Set (joinStr "" "dmg" ($gameState.Get "state")) $dmg }}
				{{ else if eq $r "🛡"}}
					{{ $actionDesc = "braces!"}}
					{{ $gameState.Set (joinStr "" "dmg" ($gameState.Get "state")) (mult $dmg -1) }}
				{{ end }}
			{{ end }}
			
			{{ $health1 := ($gameState.Get "health1")}}
			{{ $health2 := ($gameState.Get "health2")}}

			{{ $state = add $state 1 }}
			{{ $gameState.Set "state" $state }}
			
			{{ $embedFields := (cslice
				(sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true ))
			}}

			{{ if eq $state 3}}
				{{/*Game is at calculation state*/}}
				{{ $dmg1 := ($gameState.Get "dmg1")}}
				{{ $dmg2 := ($gameState.Get "dmg2")}}
				{{ $parry := (add $dmg2 $dmg1) }}

				{{$appendStr1 := "*Blocked!*"}}
				{{$appendStr2 := "*Blocked!*"}}
				

				{{ if gt $dmg1 0 }}
					{{/*User is attacking*/}}
					{{ $appendStr1 = (joinStr "" "Attacks for `" $dmg1 "`") }}
					{{ if lt $dmg2 0 }}
						{{/*Target is blocking*/}}
						{{ if le $parry 0 }}
							{{/*success*/}}	
							{{ $health1 = add $health1 (mult $parry 2) }}
							{{ $appendStr2 = (joinStr "" "Parried for `" (mult $parry -1) "`")}}
							{{ $appendStr1 = "Got parried!"}}
						{{ else }}
							{{ $appendStr2 = "Failed to block!"}}
							{{ $health2 = sub $health2 $dmg1 }}
						{{ end }}
					{{ else }}
						{{ $health2 = sub $health2 $dmg1 }}
					{{ end }}
				{{ end }}

				{{ if gt $dmg2 0 }}
					{{/*User is attacking*/}}
					{{ $appendStr2 = (joinStr "" "Attacks for `" $dmg2 "`") }}
					{{ if lt $dmg1 0 }}
						{{/*Target is blocking*/}}
						{{ if le $parry 0 }}
							{{/*success*/}}	
							{{ $health2 = add $health2 (mult $parry 2) }}
							{{ $appendStr1 = (joinStr "" "Parried for `" (mult $parry -1) "`")}}
							{{ $appendStr2 = "Got parried!"}}
						{{ else }}
							{{ $appendStr1 = "Failed to block!"}}
							{{ $health1 = sub $health1 $dmg2 }}
						{{ end }}
					{{ else }}
						{{ $health1 = sub $health1 $dmg2 }}
					{{ end }}
				{{ end }}

				{{ $embedFields = $embedFields.AppendSlice (cslice (sdict "name" $user1.Username "value" $appendStr1 "inline" true ) 
					(sdict "name" $user2.Username "value" $appendStr2 "inline" true ) 
					(sdict "name" "Results:" "value" "`After last round:`" "inline" true ))
				}}
				
				{{ $gameState.Set "health1" $health1 }}
				{{ $gameState.Set "health2" $health2 }}

				{{ deleteAllMessageReactions nil $x }}
				{{ addMessageReactions nil $x "⚔" "🛡" }}
				{{ $gameState.Set "state" 1 }}

				{{ if or (le $health1 0) (le $health2 0) }}
					{{ $state = 4 }}
					{{ $gameState.Set "state" 4 }}
				{{ end }}
			{{ else }}
				{{/*Next player's turn, reset reactions*/}}
				{{ deleteAllMessageReactions nil $x }}
				{{ addMessageReactions nil $x "⚔" "🛡" }}
			{{ end }}

			{{ $embedFields = $embedFields.AppendSlice (cslice (sdict "name" $user1.Username "value" (joinStr "" $health1 "♥" ) "inline" true ) 
				(sdict "name" $user2.Username "value" (joinStr "" $health2 "♥" ) "inline" true )
				(sdict "name" (joinStr "" $curUser.Username " " $actionDesc ) "value" (joinStr "" $opponent.Username " is up!") "inline" false ))
			}}

			{{$actionEmbed := cembed
				"title" (joinStr "" "__" $user1.Username "__ challenges __" $user2.Username "__ to a duel!")
				"description" "The game is on!"
				"fields" $embedFields 
			}}
			{{ editMessage nil $x (complexMessageEdit "content" $opponent.Mention "embed" $actionEmbed) }}
			
			{{ dbSetExpire $x $gameKey $gameState 60 }}
			{{ dbSetExpire $user1.ID $gameKey (sdict "active" 1 "messageID" $x) 60 }}
			{{ dbSetExpire $user2.ID $gameKey (sdict "active" 1 "messageID" $x) 60 }}

			{{ if eq $state 4}}
				{{/*Game is over, shut it down and hand out winnings.*/}}
				{{ deleteAllMessageReactions nil $x }}
				{{ $winFields := cslice (sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true ) }}

				{{ $user1History := (dbGet $user1.ID $historyKey).Value }}
				{{ $user2History := (dbGet $user2.ID $historyKey).Value }}
				{{if not (eq (printf "%T" $user1History) "*templates.SDict")}} {{ $user1History = (sdict "wins" 0 "losses" 0 )}} {{ end }}
				{{if not (eq (printf "%T" $user2History) "*templates.SDict")}} {{ $user2History = (sdict "wins" 0 "losses" 0 )}} {{ end }}

				{{ $user1balance := (dbGet $user1.ID $balanceKey).Value }}
				{{ $user2balance := (dbGet $user2.ID $balanceKey).Value }}

				{{ if and (le $health1 0) (le $health2 0)}}
					{{ $winFields = $winFields.Append (sdict "name" "Game Over!" "value" "It was a draw!"  "inline" true )}}
					{{ $user1balance = (add $user1balance $amount )}}
					{{ $user2balance = (add $user1balance $amount )}}
					{{ dbSet $user1.ID $balanceKey $user1balance }}
					{{ dbSet $user2.ID $balanceKey $user2balance }}
				{{ else if gt $health1 $health2}}
					{{/*User 1 wins*/}}
					{{ $winFields = $winFields.Append (sdict "name" "Game Over!" "value" (joinStr "" $user1.Username " won!")  "inline" true )}}
					{{ $user1balance = (add $user1balance (mult $amount 2))}}
					{{ dbSet $user1.ID $balanceKey $user1balance }}
					
					{{ $user1History.Set "wins" (add (toInt ($user1History.Get "wins")) 1)}}
					{{ $user2History.Set "losses" (add (toInt ($user2History.Get "losses")) 1)}}
				{{ else }}
					{{/*User 2 wins*/}}
					{{ $winFields = $winFields.Append (sdict "name" "Game Over!" "value" (joinStr "" $user2.Username " won!")  "inline" true )}}
					{{ $user2balance = (add $user2balance (mult $amount 2))}}
					{{ dbSet $user2.ID $balanceKey $user2balance }}

					{{ $user1History.Set "losses" (add (toInt ($user1History.Get "losses")) 1)}}
					{{ $user2History.Set "wins" (add (toInt ($user2History.Get "wins")) 1)}}
				{{ end }}
				
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