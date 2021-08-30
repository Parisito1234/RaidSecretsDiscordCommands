{{$key := "RSCoinBalance" }}
{{$blackjackKey := "BlackJackDatabase"}}

{{$initBalance := toInt (dbGet $.User.ID $key).Value }}
{{$e := "<:RSStonkCoin:869340420692394095>"}}
{{$deck := cslice
  "S1" "S2" "S3" "S4" "S5" "S6" "S7" "S8" "S9" "S10" "S11" "S12" "S13" "H1" "H2" "H3" "H4" "H5" "H6" "H7" "H8" "H9" "H10" "H11" "H12" "H13" "C1" "C2" "C3" "C4" "C5" "C6" "C7" "C8" "C9" "C10" "C11" "C12" "C13" "D1" "D2" "D3" "D4" "D5" "D6" "D7" "D8" "D9" "D10" "D11" "D12" "D13" }}


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
      {{ .Set "Res" (slice $data 1) }} 
    {{ end }}
{{ end }}

{{ define "check"}}
    {{$total := 0}}
    {{$list := .List}}
    {{range $index, $element := $list}}
      {{$number := (toInt (slice $element 1))}}
      {{if gt $number 10}}
        {{$number = 10}}
      {{else if eq $number 1}}
        {{$number = 11}}
      {{end}}

      {{$total = (add $total $number)}}
    {{end}}
    
    {{.Set "Tot" $total}}
    {{.Set "Ace" true}}
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
        {{$element = joinStr "" $suit $number}}
        {{$list.Set $index $element}}
    {{end}}
    {{.Set "Ret" $list}}
{{end}}

{{/*checks command is valid and sufficient balance*/}}
{{if ne (len .Args) 2}}
  {{sendMessage nil "Incorrect Syntax"}}
{{else}}
  {{$bet := toInt (index .Args 1)}}
  {{if le $bet 0}}
    {{sendMessage nil "Incorrect bet"}}
  {{else}}
    {{if gt $bet $initBalance}}
      {{sendMessage nil "Insufficient funds"}}
    {{else}}


{{/*start command content*/}}
{{$newbalance := (sub $initBalance $bet)}}
{{dbSet $.User.ID $key $newbalance}}


{{/*get cards. card1 and card2 are player cards, dealercard1 and dealercard2 are dealers cards*/}}
{{$roll1 := randInt (sub (len ($deck)) 1)}}
{{$card1 := (index $deck $roll1)}}
{{ $data := sdict "Data" $deck "Index" $roll1 }}
{{ template "remove" $data }} {{ $deck = $data.Res }}

{{$roll2 := randInt (sub (len ($deck)) 1)}}
{{$card2 := (index $deck $roll2)}}
{{ $data := sdict "Data" $deck "Index" $roll2 }}
{{ template "remove" $data }} {{ $deck = $data.Res }}

{{$roll3 := randInt (sub (len ($deck)) 1)}}
{{$dealercard1 := (index $deck $roll3)}}
{{ $data := sdict "Data" $deck "Index" $roll3 }}
{{ template "remove" $data }} {{ $deck = $data.Res }}

{{$roll4 := randInt (sub (len ($deck)) 1)}}
{{$dealercard2 := (index $deck $roll4)}}
{{ $data := sdict "Data" $deck "Index" $roll4 }}
{{ template "remove" $data }} {{ $deck = $data.Res }}

{{$playerhand := cslice $card1 $card2}}
{{$dealerhand := cslice $dealercard1 $dealercard2}}

{{ $data := sdict "List" $playerhand}}
{{template "check" $data}}{{$playertotal := $data.Tot}}{{$ace := $data.Ace}}

{{ $data := sdict "List" $playerhand}}
{{ template "convert" $data }} {{ $prettyplayerhand := $data.Ret }}
{{ $data := sdict "List" $dealerhand}}
{{ template "convert" $data }} {{ $prettydealerhand := $data.Ret }}

{{$player := ""}}
{{range $index, $element := $prettyplayerhand}}
  {{$player = joinStr "" $player " " $element " "}}
{{end}}

{{if ne $playertotal 21}}



{{$embed := cembed
  "title" (joinStr "" "__" $.User.Username "__ is at the blackjack table.")
  "description" (joinStr "" "They have placed `" $bet "` " $e " on the table... \nSweeper deals the cards...")
  "fields" (cslice
    (sdict "name" $.User.Username "value" $player "inline" true)
    (sdict "name" "**Sweeper**" "value" (joinStr "" (toString (index $prettydealerhand 0)) "  ?") "inline" true)
  )}}
{{$x := sendMessageRetID nil $embed}}
{{dbSet $.User.ID $blackjackKey (sdict "PlayerHand" $playerhand "DealerHand" $dealerhand "Ace" $ace "PlayerTotal" $playertotal "Deck" $deck)}}
{{addMessageReactions nil $x ":hit:881273087499337728" ":stand:881273109670412368"}}



{{else}}
{{/*PLAYER WINS WITH BLACKJACK*/}}
{{$embed := cembed
  "title" (joinStr "" "__" $.User.Username "__ is at the blackjack table.")
  "description" (joinStr "" "They have placed `" $bet "` " $e " on the table... \nSweeper deals the cards...")
  "fields" (cslice
    (sdict "name" $.User.Username "value" $player "inline" true)
    (sdict "name" "**Sweeper**" "value" (joinStr "" (toString (index $dealerhand 0)) "  ?") "inline" true)
    (sdict "name" "__**Nice!**__" "value" (joinStr "" "***BLACKJACK! You have won `" (mult $bet 2) "` !***") "inline" false)
  )}}
{{sendMessage nil $embed}}
{{dbSet $.User.ID $key (add $newbalance (mult $bet 2))}}
{{end}}
{{end}}{{end}}{{end}}
