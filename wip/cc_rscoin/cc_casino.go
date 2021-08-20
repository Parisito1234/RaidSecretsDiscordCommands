{{$e := "<:RSStonkCoin:869340420692394095>"}}
{{$black := (cslice 1 2 3 4 6 7 8 9)}}
{{$colour := 0}}
{{$win := 0}}
{{$types := (cslice "red" "black" "green" "odd" "even" 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20)}}
{{if eq (len .Args) 3}}
  {{$type := index .Args 1}}
  {{$amount := index .Args 2}}
  {{if in $types $type}}
    {{$database := true}}
    {{if eq $database true}}

      {{$roll := randInt 20}} {{/*generates random number and allocates colour, odd/even etc*/}}
      {{$roll_oddeven := (toInt (mod $roll 2))}}
      {{if eq (in $black $roll) true}}
        {{$colour = "black"}}
      {{else if or (eq $roll 5) (eq  $roll 10)}}
        {{$colour = "green"}}
      {{else}}
        {{$colour = "red"}}
      {{end}}

      {{if or (and (eq $type $colour) (ne $colour "green"))
              (and (eq $type "odd") (eq $roll_oddeven 1))
              (and (eq $type "even") (eq $roll_oddeven 0))}} {{/*calculates winnings for red/black/even/odd*/}}
        {{$win = (mult 2 $amount)}}
      {{else if and (eq $type $colour) (eq $colour "green")}}
        {{$win = (mult 8 $amount)}}
      {{else if eq (str $roll) $type}}
        {{$win = (mult 15 $amount)}}
      {{end}}

      roll: {{$roll}} odd:{{$roll_oddeven}} colour: {{$colour}} win: {{$win}}
    {{else}}
      Not enough {{$e}}, nerd.
    {{end}}
  {{else}}
  Valid bet types are: `\``red, black, green, odd, even, or any number from 1-20`\``
  {{end}}
{{else}}
  Nice try, but `-roulette <red|black|green|odd|even|number> <bet-amount>` is the correct syntax!
{{end}}