{{$key := "RSCoinBalance" }}
{{$blackjackKey := "RSCasinoBlackJack"}}
{{$e := "<:RSStonkCoin:869340420692394095>"}}
{{$initialData := (dbGet $.User.ID $blackjackKey).Value}}

{{ $silent := ""}}

{{ define "remove" }}
    {{ $data := .Data }} {{ $index := .Index }}
    {{ $last := sub (len $data) 1 }}
    {{ if gt $index $last }}
        {{ .Set "Res" 0 }}
    {{ else if eq $index $last }}
        {{ .Set "Res" (slice $data 0 $last) }}
    {{ else if $index }}
        {{ .Set "Res" ((slice $data 0 $index).AppendSlice (slice $data (add $index 1))) }}
    {{ else }}
        {{ .Set "Res" (slice $data 1) }} {{ end }}
{{ end }}

{{ define "check"}}
    {{$aceCount := 0}}
    {{$total := 0}}
    {{$list := .List}}
    {{range $index, $element := $list}}
      {{$number := (toInt (slice $element 1))}}
      {{if gt $number 10}}
          {{$number = 10}}
      {{else if eq $number 1}}
          {{$number = 11}}
          {{$aceCount = (add $aceCount 1)}}
      {{end}}
      {{$total = (add $total $number)}}
    {{end}}
    {{.Set "Tot" $total}}
    {{.Set "Ace" $aceCount}}
{{end}}

{{define "convert"}}
    {{$list := .List}}
    {{range $index, $element := $list}}
        {{$number := (toInt (slice $element 1))}}
        {{$suit := (slice $element 0 1)}}

        {{if eq $suit "H"}} {{$suit = ":hearts:"}}
        {{else if eq $suit "C"}} {{$suit = ":clubs:"}}
        {{else if eq $suit "S"}} {{$suit = ":spades:"}}
        {{else if eq $suit "D"}} {{$suit = ":diamonds:"}}{{end}}

         {{if eq $number 11}} {{$number = "J"}}
        {{else if eq $number 12}} {{$number = "Q"}}
        {{else if eq $number 1}} {{$number = "A"}}
        {{else if eq $number 13}} {{$number = "K"}}{{end}}

        {{$element = joinStr "" $suit $number}}
        {{$list.Set $index $element}}
    {{end}}
    {{.Set "Ret" $list}}
{{end}}

{{if eq (printf "%T" $initialData) "*templates.SDict"}}
	{{$deck := $initialData.Get "Deck"}}
	{{$dealerhand := $initialData.Get "DealerHand"}}
	{{$playerhand := $initialData.Get "PlayerHand"}}
	{{$playertotal := $initialData.Get "PlayerTotal"}}
	{{$x := $initialData.Get "X"}}
	{{$ace := $initialData.Get "Ace"}}
	{{$bet := $initialData.Get "Bet"}}

	{{if eq $x $.Reaction.MessageID}} {{/*correct person*/}}
		{{if eq $.Reaction.Emoji.ID 881273087499337728}} {{/*hit*/}}
			{{deleteMessageReaction nil $x $.User.ID ":hit:881273087499337728"}}

			{{$roll1 := randInt (sub (len ($deck)) 1)}}
			{{$card1 := (index $deck $roll1)}}
			{{ $data := sdict "Data" $deck "Index" $roll1 }}
			{{ template "remove" $data }} {{ $deck = $data.Res }}
			{{$playerhand2 := ($playerhand.Append $card1)}}
			{{$playerhand = ($playerhand.Append $card1)}}

			{{ $data := sdict "List" $playerhand}}
			{{template "check" $data}}{{$playertotal := $data.Tot}}{{$ace := $data.Ace}}

			{{ $data := sdict "List" $playerhand}}
			{{ template "convert" $data }} {{ $prettyplayerhand := $data.Ret }}

			{{$dealerhand2 := slice ($dealerhand.Append "") 0 2}}
			{{ $data := sdict "List" $dealerhand}}
			{{ template "convert" $data }} {{ $prettydealerhand := $data.Ret }}

			{{$player := ""}}
			{{range $index, $element := $prettyplayerhand}}
			{{$player = joinStr "" $player " " $element " "}}
			{{end}}

			{{if le (sub $playertotal (mult 10 $ace)) 21}}
				{{$embed := cembed
				"title" (joinStr "" "__" $.User.Username "__ is at the blackjack table.")
				"description" "Sweeper deals the cards..."
				"fields" (cslice
				(sdict "name" "Bet" "value" (joinStr "" $e " `" $bet "`")  "inline" true )
				(sdict "name" $.User.Username "value" $player "inline" true)
				(sdict "name" "Sweeper" "value" (joinStr "" (toString (index $prettydealerhand 0)) "  ??") "inline" true)
				)}}
				{{editMessage nil $x $embed}}
				{{$temp := (sdict "PlayerHand" $playerhand2 "DealerHand" $dealerhand2 "Ace" $ace "PlayerTotal" $playertotal "Deck" $deck "X" $x "Bet" $bet)}}
				{{dbSetExpire $.User.ID $blackjackKey $temp 120}}
			{{else if gt $playertotal 21}}
				{{$embed := cembed
				"title" (joinStr "" "__" $.User.Username "__ is at the blackjack table.")
				"description" "Sweeper deals the cards..."
				"fields" (cslice
				(sdict "name" "Bet" "value" (joinStr "" $e " `" $bet "`")  "inline" true )
				(sdict "name" $.User.Username "value" $player "inline" true)
				(sdict "name" "Sweeper" "value" (joinStr "" (toString (index $prettydealerhand 0)) "  ??") "inline" true)
				(sdict "name" "__**BUST!**__" "value" (joinStr "" "You've gone bust and lost all `" $bet "` " $e " on the table...") "inline" false)
				)}}
				{{editMessage nil $x $embed}}
				{{dbDel $.User.ID $blackjackKey}}
				{{deleteAllMessageReactions nil $x}}
			{{end}}

			{{else if eq $.Reaction.Emoji.ID 881273109670412368}}
			{{deleteMessageReaction nil $x $.User.ID ":stand:881273109670412368"}}
			{{deleteAllMessageReactions nil $x}}

			{{ $data := sdict "List" $playerhand}}
			{{template "check" $data}}{{$playertotal := $data.Tot}}{{$ace := $data.Ace}}
			{{ $data := sdict "List" $dealerhand}}
			{{template "check" $data}}{{$dealertotal := $data.Tot}}{{$dealerace := $data.Ace}}
			{{$state := 0}} {{/* 0 = empty, 1 = player wins, 2 = dealer wins, 3 = draw*/}}

			{{range seq 0 8}}
				{{if or (lt $dealertotal 17) (and (gt $dealertotal 21) (le (sub $dealertotal (mult 10 $dealerace)) 17))}}
					{{$roll := randInt (sub (len ($deck)) 1)}}
					{{$card := (index $deck $roll)}}
					{{ $data := sdict "Data" $deck "Index" $roll}}
					{{ template "remove" $data }} {{ $deck = $data.Res }}
					{{$dealerhand = ($dealerhand.Append $card)}}
					{{ $data := sdict "List" $dealerhand}}
					{{template "check" $data}}{{$dealertotal = $data.Tot}}{{$dealerace = $data.Ace}}
				{{end}}
			{{end}}
			{{$dealertotal = (sub $dealertotal (mult 10 $dealerace))}}
			{{range seq 0 $dealerace}}
				{{if lt $dealertotal 17}}
					{{$dealertotal = (add $dealertotal 10)}}
				{{end}}
			{{end}}

			{{if gt $playertotal 21}}
				{{$playertotal = (sub $playertotal (mult $ace 10))}}
			{{end}}

			{{if or (gt $playertotal $dealertotal) (gt $dealertotal 21)}}
				{{$state = 1}}
			{{else if lt $playertotal $dealertotal}}
				{{$state = 2}}
			{{else if eq $playertotal $dealertotal}}
				{{$state = 3}}
			{{end}}

			{{ $data := sdict "List" $playerhand}}
			{{ template "convert" $data }} {{ $prettyplayerhand := $data.Ret }}
			{{ $data := sdict "List" $dealerhand}}
			{{ template "convert" $data }} {{ $prettydealerhand := $data.Ret }}

			{{$player := ""}}
			{{range $index, $element := $prettyplayerhand}}{{$player = joinStr "" $player " " $element " "}}{{end}}

			{{range $index, $element := $dealerhand}}
				{{if gt $index 0}}
					{{$dealer2 := ""}}
					{{range $i, $el := (slice $prettydealerhand 0 (add $index 1))}}{{$dealer2 = joinStr "" $dealer2 " " $el " "}}{{end}}
					{{$embed := cembed
					"title" (joinStr "" "__" $.User.Username "__ is at the blackjack table.")
					"description" "The game is underway..."
					"fields" (cslice
					(sdict "name" "Bet" "value" (joinStr "" $e " `" $bet "`")  "inline" true )
					(sdict "name" $.User.Username "value" $player "inline" true)
					(sdict "name" "Sweeper" "value" $dealer2 "inline" true)
					)}}{{editMessage nil $x $embed}}{{sleep 3}}
				{{end}}
			{{end}}

			{{$dealer2 := ""}}
			{{range $i, $el := $prettydealerhand}}
				{{$dealer2 = joinStr "" $dealer2 " " $el " "}}
			{{end}}

			{{$embed := cembed
				"title" (joinStr "" "__" $.User.Username "__ is at the blackjack table.")
				"description" (joinStr "" "The game is underway...")
				"fields" (cslice
				(sdict "name" "Bet" "value" (joinStr "" $e " `" $bet "`")  "inline" true )
				(sdict "name" $.User.Username "value" (joinStr "" $player "\nTotal: " $playertotal) "inline" true)
				(sdict "name" "Sweeper" "value" (joinStr "" $dealer2 "\nTotal: " $dealertotal) "inline" true)
			)}}
			{{$silent = editMessage nil $x $embed}}

			{{$balance := toInt (dbGet $.User.ID $key).Value }}
			{{$endMsg := (sdict "temp" "temp")}}
			{{if eq $state 1}}
				{{dbSet $.User.ID $key (add $balance (mult $bet 2))}}
				{{$endMsg = (sdict "name" "**WIN!!**" "value" (joinStr "" "`" $bet "` " $e " has been added to your balance!\n`" $.User.Username "` now has " $e " `" (dbGet $.User.ID $key).Value "`") "inline" false) }}
			{{else if eq $state 2}}
				{{$endMsg = (sdict "name" "**LOST!!**" "value" (joinStr "" "Sweeper sweeps the money on the table into his pockets...\n`" $.User.Username "` now has " $e " `" (dbGet $.User.ID $key).Value "`" ) "inline" false) }}
			{{else if eq $state 3}}
				{{dbSet $.User.ID $key (add $balance $bet)}}
				{{$endMsg = (sdict "name" "**DRAW!!**" "value" (joinStr "" "You pick the money up from the table, and add it back to your pocket...\n`" $.User.Username "` now has " $e " `" (dbGet $.User.ID $key).Value "`") "inline" false)}}
			{{end}}

			{{$embed := cembed
				"title" (joinStr "" "__" $.User.Username "__ is at the blackjack table.")
				"description" "The game is over."
				"fields" (cslice
				(sdict "name" "Bet" "value" (joinStr "" $e " `" $bet "`")  "inline" true )
				(sdict "name" $.User.Username "value" (joinStr "" $player "\nTotal: " $playertotal) "inline" true)
				(sdict "name" "Sweeper" "value" (joinStr "" $dealer2 "\nTotal: " $dealertotal) "inline" true)
				$endMsg)
			}}
			{{$silent = editMessage nil $x $embed}}

			{{$silent = dbDel $.User.ID $blackjackKey}}

		{{end}}
	{{end}}
{{end}}