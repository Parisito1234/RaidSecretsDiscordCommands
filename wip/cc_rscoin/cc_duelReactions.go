{{ $balanceKey := "RSCoinBalance" }}
{{ $gameKey := "RSCasinoDuel" }}
{{$e := "<:RSStonkCoin:869340420692394095>"}}

{{ $user := .User }}
{{ $x := .Reaction.MessageID }}
{{ $gameState := (dbGet $x $gameKey).Value }}
{{ $r := .Reaction.Emoji.Name }}
{{ $user1 := userArg ($gameState.Get "user1")}}
{{ $user2 := userArg ($gameState.Get "user2")}}
{{ $amount := $gameState.Get "bet" }}

{{ if eq "pending" ($gameState.Get "state")}}
{{/* Waiting for game to start */}}

	{{ if eq $user.ID $user2.ID }}
		{{ deleteAllMessageReactions nil $x }}
		{{ if eq $r "✅"}}
			{{ $gameState.Set "health1" 100 }}
			{{ $gameState.Set "health2" 100 }}
			{{ $gameState.Set "state" "active" }}
			{{ $gameState.Set "turn" 2 }}
			{{ dbSetExpire $x $gameKey $gameState 120}}

			{{$embed := cembed
				"title" (joinStr "" "__" $user1.Username "__ challenges __" $user2.Username "__ to a duel!")
				"description" "The game is starting!"
				"fields" (cslice
				(sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true )
				(sdict "name" $user1.Username "value" (joinStr "" ($gameState.Get "health1") "♥" ) "inline" true )
				(sdict "name" $user2.Username "value" (joinStr "" ($gameState.Get "health2") "♥" )  "inline" true ))
			}}
			{{ editMessage nil $x (complexMessageEdit "content" " " "embed" $embed) }}
			{{ sleep 3 }}
			{{$embed = cembed
				"title" (joinStr "" "__" $user1.Username "__ challenges __" $user2.Username "__ to a duel!")
				"description" "The game is on!"
				"fields" (cslice
				(sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true )
				(sdict "name" $user1.Username "value" (joinStr "" ($gameState.Get "health1") "♥" ) "inline" true )
				(sdict "name" $user2.Username "value" (joinStr "" ($gameState.Get "health2") "♥" )  "inline" true ))
			}}
			{{ editMessage nil $x (complexMessageEdit "content" $user2.Mention "embed" $embed) }}
			{{ addMessageReactions nil $x ":crossed_swords:" ":shield:" }}
		{{ else if eq $r "❌" }}
			{{$embed := cembed
				"title" (joinStr "" "__" $user1.Username "__ challenged __" $user2.Username "__ to a duel.")
				"description" "The duel was declined."
			}}
			{{ dbDel $user1.ID $gameKey }}
			{{ dbDel $user2.ID $gameKey }}
			{{ dbDel $x $gameState }}
			{{ editMessage nil $x $embed }}
		{{ end }}
		
	{{ else if eq $user.ID $user1.ID }}
		{{ deleteMessageReaction nil $x .Reaction.UserID $r }}
		{{ sendMessage nil "Wait for your opponent to start the match!"}}
	{{ end }}

{{ else if eq "active" ($gameState.Get "state")}}
{{/*Game has started */}}
	{{ $curUser := $user1 }}
	{{ if eq ($gameState.Get "turn") 2}} {{ $curUser = $user2 }} {{ end }}

{{ end }}