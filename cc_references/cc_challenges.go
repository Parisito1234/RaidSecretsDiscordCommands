{{$argslength := len .CmdArgs}}
{{$args := (joinStr " " .CmdArgs) }}
{{if ge $argslength 1}}
	{{$firstarg := (index .CmdArgs 0)}}
	{{if eq $firstarg "1"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_1").Value)}}
	{{else if eq $firstarg "2"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_2").Value)}}
	{{else if eq $firstarg "3"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_3").Value)}}
	{{else if eq $firstarg "4"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_4").Value)}}
	{{else if eq $firstarg "5"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_5").Value)}}
	{{else if eq $firstarg "6"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_6").Value)}}
	{{else if eq $firstarg "7"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_7").Value)}}
	{{else if eq $firstarg "8"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_8").Value)}}
	{{else if eq $firstarg "9"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_9").Value)}}
	{{else if eq $firstarg "10"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_10").Value)}}
	{{else if eq $firstarg "all"}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_1").Value)}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_2").Value)}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_3").Value)}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_4").Value)}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_5").Value)}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_6").Value)}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_7").Value)}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_8").Value)}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_9").Value)}}
		{{sendMessage nil (cembed (dbGet 0 "reference_challenges_10").Value)}}
	{{else}}
		Please provide a valid week number for the season. This season has 10 weeks.
	{{end}}
{{else}}
	Please provide a valid query. Usage: `-challenges <week number>`
{{end}}