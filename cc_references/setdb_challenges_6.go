{{ $challenge_1 := sdict 
	"title" "Contender's Ascent VI" 
	"description" "Report to the War Table in the H.E.L.M. and complete the \"Challenger's Proving VI\" quest. Then defeat (60) Cabal using a Bow anywhere in the system." 
	"reward" "•XP  •Imperial Needle  •Empress Lore  •War Table Reputation (Large)" 
}}
{{ $challenge_2 := sdict 
	"title" "Small-Caliber Contender" 
	"description" "Defeat combatants using a Submachine Gun. Battlegrounds missions grant the most efficient progress. (400)" 
	"reward" "•XP" 
}}
{{ $challenge_3 := sdict 
	"title" "Nightfall Variety Attack " 
	"description" "Defeat (80) combatants using Sniper Rifles, Submachine Guns, Rocket Launchers, or Bows in Nightfall: The Ordeal." 
	"reward" "•XP" 
}}
{{ $challenge_4 := sdict 
	"title" "Cabal Contenders" 
	"description" "Defeat (65) Elite or Boss Cabal anywhere in Battlegrounds missions." 
	"reward" "•XP  •War Table Reputation (Large)" 
}}
{{ $challenge_5 := sdict 
	"title" "Once Chosen, Now Fallen" 
	"description" "Calibrate weapons against Fallen. Earn bonus progress for precision final blows. (200)" 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_6 := sdict 
	"title" "Kill Them with Style" 
	"description" "Acquire the Pyretic ornament for the Salvager's Salvo Grenade Launcher." 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_7 := sdict 
	"title" "Conquest of the Mighty" 
	"description" "Complete any Nightfall: The Ordeal strike on Grandmaster." 
	"reward" "•4x XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_6" (sdict 
	"title" "Seasonal Challenges - Week 6" 
	"color" 1772743
	"fields" (cslice 
		(sdict "name" ($challenge_1.Get "title") "value" (joinStr "" ($challenge_1.Get "description") " **|** `" ($challenge_1.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_2.Get "title") "value" (joinStr "" ($challenge_2.Get "description") " **|** `" ($challenge_2.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_3.Get "title") "value" (joinStr "" ($challenge_3.Get "description") " **|** `" ($challenge_3.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_4.Get "title") "value" (joinStr "" ($challenge_4.Get "description") " **|** `" ($challenge_4.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_5.Get "title") "value" (joinStr "" ($challenge_5.Get "description") " **|** `" ($challenge_5.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_6.Get "title") "value" (joinStr "" ($challenge_6.Get "description") " **|** `" ($challenge_6.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_7.Get "title") "value" (joinStr "" ($challenge_7.Get "description") " **|** `" ($challenge_7.Get "reward") "`") "inline" false) 
	) 
)}}
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_6").Value)}}