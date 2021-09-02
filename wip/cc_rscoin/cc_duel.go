{{ $balanceKey := "RSCoinBalance" }}
{{ $gameKey := "RSCasinoDuel" }}
{{ $historyKey := "RSCasinoDuelHistory" }}
{{$e := "<:RSStonkCoin:869340420692394095>"}}

{{ $args := parseArgs 2 "Syntax is `-duel <user> <amount>`"
	(carg "userid" "user to perform action on")
	(carg "int" "amount") 
}}

{{/* Dynamic data init */}}
{{ $user1 := $.User }}
{{ $user2 := (userArg ($args.Get 0))}}
{{ $user1History := (dbGet $user1.ID $historyKey).Value }}
{{ $user2History := (dbGet $user2.ID $historyKey).Value }}
{{ $amount := ($args.Get 1)}}
{{ $curBalance := toInt (dbGet $user1.ID $balanceKey).Value}}

{{/* Check if user currently has an active duel*/}}
{{ $gameState1 := (dbGet $user1.ID $gameKey).Value }}
{{ $gameState2 := (dbGet $user2.ID $gameKey).Value }}

{{ if or (eq (toInt ($gameState1.Get "active")) 1 ) (eq (toInt ($gameState2.Get "active")) 1 ) }}
	{{ sendMessage nil "One of the players still has an active duel!" }}
{{ else if eq $user1.ID $user2.ID }}
	{{ sendMessage nil (joinStr "" $user1.Username ", you can't start a duel with yourself.")}}
{{ else if le $amount 0 }}
	{{ sendMessage nil (joinStr "" $user1.Username ", your bet needs to be a positive number!")}}
{{ else if gt $amount $curBalance }}
	{{ sendMessage nil (joinStr "" $user1.Username ", you don't have enough!")}}
{{ else if gt $amount (toInt (dbGet $user2.ID $balanceKey).Value) }}
	{{ sendMessage nil (joinStr "" $user1.Username ", your opponent doesn't have enough to cover their wager!")}}
{{ else }}
	{{ deleteTrigger 0 }}

	{{if gt $amount 500}} {{$amount = 500}} {{ end }}

	{{/* Message construction and send */}}
	{{$embed := cembed
		"title" (joinStr "" "__" $user1.Username "__ challenges __" $user2.Username "__ to a duel!")
		"description" "Waiting for a response..."
		"fields" (cslice
		(sdict "name" "Wager" "value" (joinStr "" $e " `" $amount "`")  "inline" true )
		(sdict "name" $user1.Username "value" (joinStr "" "Wins: " (toInt ($user1History.Get "wins")) "\nLosses: " (toInt ($user1History.Get "losses")) ) "inline" true )
		(sdict "name" $user2.Username "value" (joinStr "" "Wins: " (toInt ($user2History.Get "wins")) "\nLosses: " (toInt ($user2History.Get "losses")) )  "inline" true ))
		"footer" (sdict "text" (joinStr "" $user2.Username " needs to react within 1 minute to continue."))
	}}

	{{ $x := sendMessageRetID nil (complexMessage "content" $user2.Mention "embed" $embed) }}

	{{ dbSetExpire $user1.ID $gameKey (sdict "active" 1 "messageID" $x) 60 }}
	{{ dbSetExpire $user2.ID $gameKey (sdict "active" 1 "messageID" $x) 60 }}
	{{ dbSetExpire $x $gameKey (sdict "state" 0 "user1" $user1.ID "user2" $user2.ID "bet" $amount) 60 }}

	{{ addMessageReactions nil $x ":white_check_mark:" ":x:" }}
{{ end }}