{{ $balanceKey := "RSCoinBalance" }}
{{ $duelKey := "RSCasinoDuel" }}
{{$e := "<:RSStonkCoin:869340420692394095>"}}

{{ $args := parseArgs 2 "Syntax is `-duel <user> <amount>`"
	(carg "userid" "user to perform action on")
	(carg "int" "amount") }}