{{ $challenge_1 := sdict 
	"title" "Contender's Ascent VII" 
	"description" "Report to the War Table in the H.E.L.M. and complete the \"Challenger's Proving VII\" quest. Then, defeat (100) Cabal with a Submachine Gun." 
	"reward" "•XP  •Extraordinary Rendition  •Empress Lore  •War Table Reputation (Large)" 
}}
{{ $challenge_2 := sdict 
	"title" "Proving Grounds Trifecta" 
	"description" "Complete the following in the Proving Grounds strike: finish the mission, defeat (150) combatants, and deposit a Power Core in the undercarriage of the Land Tank." 
	"reward" "•Double XP  •War Table Reputation (Large)" 
}}
{{ $challenge_3 := sdict 
	"title" "Diplomacy or Death" 
	"description" "Listen to intercepted transmissions at the radio kiosk in the H.E.L.M." 
	"reward" "•XP" 
}}
{{ $challenge_4 := sdict 
	"title" "One Against Many" 
	"description" "Calibrate weapons by rapidly defeating 3 or more combatants. (10)" 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_5 := sdict 
	"title" "Gambit Salvager's Salvo" 
	"description" "Acquire the Toxicology ornament for the Salvager's Salvo Grenade Launcher." 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_6 := sdict 
	"title" "Beyond Legendary " 
	"description" "Earn (5) Valor ranks." 
	"reward" "•4x XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_7" (sdict 
	"title" "Seasonal Challenges - Week 7" 
	"color" 1772743
	"fields" (cslice 
		(sdict "name" ($challenge_1.Get "title") "value" (joinStr "" ($challenge_1.Get "description") " **|** `" ($challenge_1.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_2.Get "title") "value" (joinStr "" ($challenge_2.Get "description") " **|** `" ($challenge_2.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_3.Get "title") "value" (joinStr "" ($challenge_3.Get "description") " **|** `" ($challenge_3.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_4.Get "title") "value" (joinStr "" ($challenge_4.Get "description") " **|** `" ($challenge_4.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_5.Get "title") "value" (joinStr "" ($challenge_5.Get "description") " **|** `" ($challenge_5.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_6.Get "title") "value" (joinStr "" ($challenge_6.Get "description") " **|** `" ($challenge_6.Get "reward") "`") "inline" false) 
	) 
)}}
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_7").Value)}}