{{$perms := "ManageServer"}}
{{$e := "<:RSStonkCoin:869340420692394095>"}}
{{$key := "RSCoinBalance" }}
{{ $lotteryPool := toInt ((dbGet 204255221017214977 $key).Value) }}
{{$roll := randInt 2}}
{{$rollType := 0}}
{{$winState := 0}}
{{$curBalance := toInt (dbGet $.User.ID $key).Value}}
{{if eq (len .Args) 3}}

  {{$amount := toInt (index .Args 2)}}
  {{$bet := lower (index .Args 1)}}
  {{if lt $amount $curBalance}}
  
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
      {{$winState = 2}}                                 
    {{end}}
    
    {{if eq $roll 1}}
      {{$rollType = "Heads"}}
    {{else}}
      {{$rollType = "Tails"}}
    {{end}}
    
    {{if eq $winState 1}}
      {{sendMessage nil (joinString "Congrats " .User.Username " has won `" $amount "`!!!")}}
    
    
    
    
    
  {{else}}
    {{sendMessage nil (joinString "Insufficient " $e)}}
  {{end}}
{{else}}
  {{sendMessage nil "Incorrect Syntax"}}
{{end}}
