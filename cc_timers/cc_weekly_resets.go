{{ /* Trigger: Hourly interval, excluding all hours except for 17:00 UTC */}}
{{ /* Cleans daily messages every day, cleans all messages every week. */}}
{{ /* Channel: #reset-info or equivalent */ }}

{{ $weekday := str currentTime.Weekday }}
{{ if eq $weekday "Tuesday" }}
	{{ exec "clean -nopin -minage 3h 20" }}
{{ else }}
	{{ exec "clean -nopin -ma 25h -minage 12h <@296023718839451649> 15" }}
	{{ exec "clean -nopin -r Reset -i 5" }}
{{ end }}
{{ sendMessage nil (joinStr "" "Reset info for **" currentTime.Weekday ", " currentTime.Month " " currentTime.Day ":**" )}}