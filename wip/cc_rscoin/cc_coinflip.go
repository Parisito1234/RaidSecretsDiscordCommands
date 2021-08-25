{{$perms := "ManageServer"}}
{{$e := "<:RSStonkCoin:869340420692394095>"}}
{{$key := "RSCoinBalance" }}
{{ $lotteryPool := toInt ((dbGet 204255221017214977 $key).Value) }}
{{$roll := randInt 2}}
{{$rollType := 0}}
{{$curBalance := toInt (dbGet $.User.ID $key).Value}}
{{if eq (len .Args) 3}}

  {{$amount := toInt (index .Args 2)}}
  {{if gt $amount 50}} {{$amount = 50}} {{sendMessage nil (joinStr "" "You can only bet up to 50 " $e " with coinflip")}} {{end}}
  {{$bet := lower (index .Args 1)}}
  {{if and (lt $amount $curBalance) (ne $amount 0)}}
    {{$winState := 2}}

    {{if or (eq $bet "head") (eq $bet "heads")}}
      {{if eq $roll 1}}
        {{$winState  = 1}}
      {{end}}
    {{else if or (eq $bet "tails") (eq $bet "tail")}}
      {{if eq $roll 2}}
        {{$winState = 1}}
      {{end}}
    {{else}}
      {{sendMessage nil "Its literally only heads or tails bruh"}}
      {{$winState = 3}}                                 
    {{end}}

    {{if eq $roll 1}}
      {{$rollType = "heads"}}
    {{else}}
      {{$rollType = "tails"}}
    {{end}}

    {{if eq $winState 1}}
      {{sendMessage nil (joinStr "" "The coin landed on " $rollType " and " .User.Username " has won `" $amount "`!!!\n" $.User.Username " now has `" (add $curBalance $amount) "` " $e)}}
      {{dbSet $.User.ID $key (toString (add $curBalance $amount))}}
    {{else if eq $winState 2}}
      {{sendMessage nil (joinStr "" "Nope! It was " $rollType "!")}}
      {{execAdmin "thrallpit" $.User (joinStr "" $amount "h CoinFlips")}}
    {{end}}

  {{else}}
    {{sendMessage nil (joinStr "" "Insufficient " $e ", or invalid amount.")}}
  {{end}}
{{else}}
  {{sendMessage nil "Incorrect Syntax"}}
{{end}}