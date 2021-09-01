{{$perms := "ManageServer"}}
{{$e := "<:RSStonkCoin:869340420692394095>"}}
{{$black := (cslice 1 2 3 4 6 7 8 9)}}
{{$color := 0}}
{{$win := 0}}
{{$key := "RSCoinBalance" }}
{{ $lotteryPool := toInt ((dbGet 204255221017214977 $key).Value) }}
{{$opts := (cslice "red" "black" "green" "odd" "even" "1" "2" "3" "4" "5" "6" "7" "8" "9" "10" "11" "12" "13" "14" "15" "16" "17" "18")}}
{{$imgur := sdict
"1" "fDtHkij"
"2" "CPUvkRj"
"3" "drgbCf3"
"4" "9SsM5ju"
"5" "foJhwxa"
"6" "5ZXgbTl"
"7" "VhGaKxB"
"8" "LRSjXKK"
"9" "8mD7DSZ"
"10" "rpHQ2YW"
"11" "kaaspHO"
"12" "P21TYin"
"13" "xSUvKex"
"14" "D5bGDzt"
"15" "CiDpI7g"
"16" "ubWRqXi"
"17" "9yiWKXV"
"18" "FyWxGyM"
}}

{{$source := "https://i.imgur.com/PKryC4s.png"}}
{{if eq (len .Args) 3}}
  {{$opt := lower (index .Args 1)}}
  {{$amount := toInt (index .Args 2)}}

  {{if in $opts $opt }}
    {{ $curBalance := toInt (dbGet $.User.ID $key).Value }}

    {{if and (ge $curBalance $amount) (gt $amount 0)}}
      {{ $embed := cembed "title" (joinStr "" "__" .User.Username "__ is at the roulette table") "description" (joinStr "" "They have bet `" (toString $amount) "` that the ball will land on `" $opt "`")  "thumbnail" (sdict "url" $source)}}

      {{$x := sendMessageRetID nil $embed }}
      {{sleep 4}}
      {{$roll := add 1 (randInt 18)}} {{/*rng allocates color, odd/even etc*/}}
      {{$roll_oddeven := (toInt (mod $roll 2))}}

      {{if eq (in $black $roll) true}}
        {{$color = "black"}}
      {{else if or (eq $roll 5) (eq  $roll 10)}}
        {{$color = "green"}}
      {{else}}
        {{$color = "red"}}
      {{end}}

      {{if or (and (eq $opt $color) (ne $color "green"))
              (and (eq $opt "odd") (eq $roll_oddeven 1))
              (and (eq $opt "even") (eq $roll_oddeven 0))}} {{/*calculates winnings*/}}
        {{$win = (mult 1 $amount)}}
      {{else if and (eq $opt $color) (eq $color "green")}}
        {{$win = (mult 7 $amount)}}
      {{else if eq (str $roll) $opt}}
        {{$win = (mult 14 $amount)}}
      {{end}}
      {{$status := "won"}}

      {{if eq $win 0}}
        {{$win = $amount}}
        {{$status = "lost"}}
        {{if not (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	        {{dbSet 204255221017214977 $key (add $lotteryPool $amount) }}
        {{end}}
      {{end}}

      {{$value := 0}}

      {{if eq $status "lost"}}
        {{$value = sub $curBalance $amount}}
      {{else}}

        {{if lt (add $curBalance $win) 0}}
          {{$value = $curBalance}}
          {{$status = "Oops, something happened. Try again."}}
        {{else}}
          {{$value = add $curBalance $win}}
        {{end}}

      {{end}}

      {{ dbSet $.User.ID $key  (str $value) }}
      {{$link := (joinStr "" "https://i.imgur.com/" ($imgur.Get (toString $roll)) ".png" )}}
      {{$img := (sdict "url" $link)}}
      {{ $uname := .User.Username }}
      {{ $bal := (toInt (dbGet $.User.ID $key).Value) }}
      {{$embed2 := cembed
        "title" (joinStr "" $uname " has " $status " `" $win "` " $e " at the roulette table")
        "description" (joinStr "" "The ball landed on " $roll "\n\n" $uname " now has `" $bal "`" $e)
        "thumbnail" $img }}

      {{if eq $status "lost"}}
        {{$embed2 = cembed
          "title" (joinStr "" $uname " has " $status " `" $win "` " $e " at the roulette table")
          "description" (joinStr "" ":x: The ball landed on " $roll "\n\n" $uname " now has `" $bal "`" $e " \n *Their loss has gone to the lottery pool.*")
          "thumbnail" $img }}
      {{end}}

      {{editMessage nil $x $embed2}}
    {{else}}
      Not enough {{$e}}!
    {{end}}

  {{else}}
    Valid bet types are: `red, black, green, odd, even, or any number from 1-18`
  {{end}}

{{else}}
  Syntax: `-roulette <red|black|green|odd|even|number> <bet-amount>`
{{end}}
