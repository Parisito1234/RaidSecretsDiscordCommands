{{ $challenge_1 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}
{{ $challenge_2 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}
{{ $challenge_3 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}
{{ $challenge_4 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}
{{ $challenge_5 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}
{{ $challenge_6 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}
{{ $challenge_7 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}
{{ $challenge_8 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}
{{ $challenge_9 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}
{{ $challenge_10 := sdict 
	"title" "" 
	"description" "" 
	"reward" "" 
}}

{{dbSet 0 "reference_challenges_1" (sdict 
	"title" "Seasonal Challenges - Week 1" 
	"color" 1772743
	"fields" (cslice 
		(sdict "name" ($challenge_1.Get "title") "value" (joinStr "" ($challenge_1.Get "description") "\n*" ($challenge_1.Get "reward") "*") "inline" false) 
		(sdict "name" ($challenge_2.Get "title") "value" (joinStr "" ($challenge_2.Get "description") "\n*" ($challenge_2.Get "reward") "*") "inline" false) 
		(sdict "name" ($challenge_3.Get "title") "value" (joinStr "" ($challenge_3.Get "description") "\n*" ($challenge_3.Get "reward") "*") "inline" false) 
		(sdict "name" ($challenge_4.Get "title") "value" (joinStr "" ($challenge_4.Get "description") "\n*" ($challenge_4.Get "reward") "*") "inline" false) 
		(sdict "name" ($challenge_5.Get "title") "value" (joinStr "" ($challenge_5.Get "description") "\n*" ($challenge_5.Get "reward") "*") "inline" false) 
		(sdict "name" ($challenge_6.Get "title") "value" (joinStr "" ($challenge_6.Get "description") "\n*" ($challenge_6.Get "reward") "*") "inline" false) 
		(sdict "name" ($challenge_7.Get "title") "value" (joinStr "" ($challenge_7.Get "description") "\n*" ($challenge_7.Get "reward") "*") "inline" false) 
		(sdict "name" ($challenge_8.Get "title") "value" (joinStr "" ($challenge_8.Get "description") "\n*" ($challenge_8.Get "reward") "*") "inline" false) 
		(sdict "name" ($challenge_9.Get "title") "value" (joinStr "" ($challenge_9.Get "description") "\n*" ($challenge_9.Get "reward") "*") "inline" false) 
		(sdict "name" ($challenge_10.Get "title") "value" (joinStr "" ($challenge_10.Get "description") "\n*" ($challenge_10.Get "reward") "*") "inline" false) 
	) 
)}}
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_1").Value)}}