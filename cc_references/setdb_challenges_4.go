{{ $challenge_1 := sdict 
	"title" "Contender's Ascent IV" 
	"description" "Report to the War Table in the H.E.L.M. and complete the \"Challenger's Proving IV\" quest. Then defeat (50) Cabal using Sniper Rifles." 
	"reward" "•XP  •Far Future  •Empress Lore  •War Table Reputation (Medium)" 
}}
{{ $challenge_2 := sdict 
	"title" "Golden Reaper" 
	"description" "Acquire (400) Cabal Gold by playing strikes, Gambit, Crucible, public events, and more." 
	"reward" "•XP  •War Table Reputation (Medium)" 
}}
{{ $challenge_3 := sdict 
	"title" "Intruder Alert" 
	"description" "Don't allow opponents to interrupt Ghost (2) in Battlegrounds missions." 
	"reward" "•XP" 
}}
{{ $challenge_4 := sdict 
	"title" "Explosive Entrance" 
	"description" "Defeat (10) Champions and get (45) grenade takedowns in Battlegrounds." 
	"reward" "•XP  •War Table Reputation (Medium)" 
}}
{{ $challenge_5 := sdict 
	"title" "Contender's Delve" 
	"description" "Complete a Lost Sector on Legend or higher." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_6 := sdict 
	"title" "Challenger's Cipher" 
	"description" "Decrypt (5) Prime Engrams." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_7 := sdict 
	"title" "Sling the Stone…" 
	"description" "Stagger, pierce, or disrupt (30) Champions." 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_8 := sdict 
	"title" "High-Value Hunter" 
	"description" "Defeat (75) powerful combatants in Gambit. Earn bonus progress for defeating high-value targets." 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_9 := sdict 
	"title" "Cadre of Contenders" 
	"description" "Complete (15) Crucible matches in the Competitive playlist." 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_10 := sdict 
	"title" "Vanguard Chosen" 
	"description" "Complete any (3) Nightfall: The Ordeal strike on Hero difficulty or higher." 
	"reward" "•XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_4" (sdict 
	"title" "Seasonal Challenges - Week 4" 
	"color" 1772743
	"fields" (cslice 
		(sdict "name" ($challenge_1.Get "title") "value" (joinStr "" ($challenge_1.Get "description") " **|** `" ($challenge_1.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_2.Get "title") "value" (joinStr "" ($challenge_2.Get "description") " **|** `" ($challenge_2.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_3.Get "title") "value" (joinStr "" ($challenge_3.Get "description") " **|** `" ($challenge_3.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_4.Get "title") "value" (joinStr "" ($challenge_4.Get "description") " **|** `" ($challenge_4.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_5.Get "title") "value" (joinStr "" ($challenge_5.Get "description") " **|** `" ($challenge_5.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_6.Get "title") "value" (joinStr "" ($challenge_6.Get "description") " **|** `" ($challenge_6.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_7.Get "title") "value" (joinStr "" ($challenge_7.Get "description") " **|** `" ($challenge_7.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_8.Get "title") "value" (joinStr "" ($challenge_8.Get "description") " **|** `" ($challenge_8.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_9.Get "title") "value" (joinStr "" ($challenge_9.Get "description") " **|** `" ($challenge_9.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_10.Get "title") "value" (joinStr "" ($challenge_10.Get "description") " **|** `" ($challenge_10.Get "reward") "`") "inline" false) 
	) 
)}}
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_4").Value)}}