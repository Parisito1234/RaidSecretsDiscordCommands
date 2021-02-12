{{ $challenge_1 := sdict 
	"title" "Convex Convergence" 
	"description" "Focus (1) Tier 3 Umbral Engrams." 
	"reward" "•XP" 
}}
{{ $challenge_2 := sdict 
	"title" "Suited for Combat" 
	"description" "Defeat targets anywhere in the system with a full armor set. Bonus progress for defeating Guardians. (400)" 
	"reward" "•Double XP" 
}}
{{ $challenge_3 := sdict 
	"title" "<Classified> Dire Portents" 
	"description" "Complete Exotic quest \"Presage\" on Normal or Master difficulty." 
	"reward" "•Classified" 
}}
{{ $challenge_4 := sdict 
	"title" "…And Fell the Giant" 
	"description" "Stagger, pierce, or disrupt (50) Champions." 
	"reward" "•4x XP  •Bright Dust" 
}}
{{ $challenge_5 := sdict 
	"title" "Feels Good to Be Bad" 
	"description" "Defeat (40) Guardians in Gambit. Earn bonus progress for defeating them as an Invader." 
	"reward" "•Double XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_9" (sdict 
	"title" "Seasonal Challenges - Week 9" 
	"color" 1772743
	"fields" (cslice 
		(sdict "name" ($challenge_1.Get "title") "value" (joinStr "" ($challenge_1.Get "description") " **|** `" ($challenge_1.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_2.Get "title") "value" (joinStr "" ($challenge_2.Get "description") " **|** `" ($challenge_2.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_3.Get "title") "value" (joinStr "" ($challenge_3.Get "description") " **|** `" ($challenge_3.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_4.Get "title") "value" (joinStr "" ($challenge_4.Get "description") " **|** `" ($challenge_4.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_5.Get "title") "value" (joinStr "" ($challenge_5.Get "description") " **|** `" ($challenge_5.Get "reward") "`") "inline" false) 
	) 
)}}
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_9").Value)}}