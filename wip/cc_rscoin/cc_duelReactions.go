{{ $balanceKey := "RSCoinBalance" }}
{{ $gameKey := "RSCasinoDuel" }}
{{$e := "<:RSStonkCoin:869340420692394095>"}}

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
				{{ dbSetExpire $x $gameKey $gameState 120}}

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
				{{ editMessage nil $x (complexMessageEdit "content" $user2.Mention "embed" $embed) }}
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

			{{$appendStr1 := "*Blocked!*"}}
			{{$appendStr2 := "*Blocked!*"}}

			{{ if gt $dmg1 0 }}
				{{ $appendStr1 = (joinStr "" "Attacks for `" $dmg1 "`") }}
				{{ if gt (add $dmg1 $dmg2) 0}}
					{{ $health2 = sub $health2 $dmg1 }}
					{{ if lt $dmg2 0 }}
						{{ $appendStr2 = "*Failed to block!*"}}
					{{ end }}
				{{ end }}
			{{ end }}

			{{ if gt $dmg2 0 }}
				{{ $appendStr2 = (joinStr "" "Attacks for `" $dmg2 "`") }}
				{{ if gt (add $dmg1 $dmg2) 0}}
					{{ $health1 = sub $health1 $dmg2 }}
					{{ if lt $dmg1 0 }}
						{{ $appendStr1 = "*Failed to block!*"}}
					{{ end }}
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

		{{ if eq $state 4}}
			{{ deleteAllMessageReactions nil $x }}
			{{ $winFields := cslice (sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true ) }}

			{{ if gt $health1 $health2}}
				{{ $winFields = $winFields.Append (sdict "name" $user1.Username "value" "won!"  "inline" true )}}
			{{ else }}
				{{ $winFields = $winFields.Append (sdict "name" $user2.Username "value" "won!"  "inline" true )}}
			{{ end }}

			{{$embed := cembed
				"title" (joinStr "" "__" $user1.Username "__ challenged __" $user2.Username "__ to a duel!")
				"description" "__Game Over!__"
				"fields" $winFields
			}}
			{{ editMessage nil $x (complexMessageEdit "content" " " "embed" $embed) }}
			
			{{ dbDel $user1.ID $gameKey }}
			{{ dbDel $user2.ID $gameKey }}
			{{ dbDel $x $gameKey }}

		{{ end }}

	{{ end }}
{{ end }}