{{$e := "<:RSStonkCoin:869340420692394095>"}}
{{$black := (cslice 1 2 3 4 6 7 8 9)}}
{{$colour := 0}}
{{$win := 0}}
{{$key := "RSCoinBalance" }}
{{$tadBalance := toInt (dbGet 255025368149393408 $key).Value }}
{{$types := (cslice "red" "black" "green" "odd" "even" "1" "2" "3" "4" "5" "6" "7" "8" "9" "10" "11" "12" "13" "14" "15" "16" "17" "18")}}
{{$imgur := sdict
"1" "https://i.imgur.com/fDtHkij.png"
"2" "https://i.imgur.com/CPUvkRj.png"
"3" "https://i.imgur.com/drgbCf3.png"
"4" "https://i.imgur.com/9SsM5ju.png"
"5" "https://i.imgur.com/foJhwxa.png"
"6" "https://i.imgur.com/5ZXgbTl.png"
"7" "https://i.imgur.com/VhGaKxB.png"
"8" "https://i.imgur.com/LRSjXKK.png"
"9" "https://i.imgur.com/8mD7DSZ.png"
"10" "https://i.imgur.com/rpHQ2YW.png"
"11" "https://i.imgur.com/kaaspHO.png"
"12" "https://i.imgur.com/P21TYin.png"
"13" "https://i.imgur.com/xSUvKex.png"
"14" "https://i.imgur.com/D5bGDzt.png"
"15" "https://i.imgur.com/CiDpI7g.png"
"16" "https://i.imgur.com/ubWRqXi.png"
"17" "https://i.imgur.com/9yiWKXV.png"
"18" "https://i.imgur.com/FyWxGyM.png"

  }}

{{$source := "https://i.imgur.com/PKryC4s.png"}}
{{if eq (len .Args) 3}}
  {{$type := index .Args 1}}
  {{$amount := toInt (index .Args 2)}}
  {{if in $types $type}}
    {{ $curBalance := toInt (dbGet $.User.ID $key).Value }}
    {{if and (ge $curBalance $amount) (gt $amount 0)}}
      {{ $embed := cembed "title" (joinStr "" .User.Username " is at the roulette table") "description" (joinStr "" "They have bet `" (toString $amount) "` that the ball will land on `" $type "`.") "image" (sdict "url" $source)}}

      {{$x := sendMessageRetID nil $embed }}
      {{sleep 5}}
      {{$roll := randInt 18}} {{/*generates random number and allocates colour, odd/even etc*/}}
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
      {{$status := "won"}}
      {{$win2 := 0}}
      {{if eq $win 0}}
        {{$win = $amount}}
        {{$status = "lost"}}
        {{$win2 = $win}}
	{{dbSet 255025368149393408 $key  (str (add $tadBalance $amount)) }}
      {{else}}
        {{$win2 = (sub $win $amount)}}
      {{end}}

      {{$value := 0}}
      {{if eq $status "lost"}}
        {{$value = sub $curBalance $amount}}
      {{else}}
        {{if lt (add $curBalance $win) 0}}
          {{$value = $curBalance}}
          {{$status = "BROKEN INTEGER OVERFLOW, WE DIDNT TOUCH YOUR BALANCE"}}
        {{else}}
          {{$value = add $curBalance (sub $win $amount)}}
        {{end}}
      {{end}}
      {{ dbSet $.User.ID $key  (str $value) }}
      {{$link := $imgur.Get (toString $roll)}}
      {{$img := (sdict "url" $link)}}
      {{$embed2 := cembed
        "title" (joinStr "" .User.Username " has " $status " `" $win2 "` " $e " at the roulette table")
        "description" (joinStr "" "The ball landed on " $roll ".\n\n"
          .User.Username " now has " (toInt (dbGet $.User.ID $key).Value) $e)
        "image" $img}}
      {{editMessage nil $x $embed2}}
    {{else}}
      Not enough {{$e}}, nerd.
    {{end}}
  {{else}}
  Valid bet types are: ```red, black, green, odd, even, or any number from 1-18```
  {{end}}
{{else}}
  Nice try, but `-roulette <red|black|green|odd|even|number> <bet-amount>` is the correct syntax!
{{end}}
