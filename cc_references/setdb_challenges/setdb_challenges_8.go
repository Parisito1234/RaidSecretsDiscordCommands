{{ $challenge_1 := sdict 
	"title" "Nightfall Proving Grounds" 
	"description" "Defeat (50) combatants using a Linear Fusion Rifle or Sidearm in (3) the Nightfall: The Ordeal version of the Proving Grounds strike." 
	"reward" "•XP" 
}}
{{ $challenge_2 := sdict 
	"title" "Legendary Lost Sector Variety Attack" 
	"description" "Defeat combatants in Legendary Lost Sectors using Sniper Rifles, Submachine Guns, Rocket Launchers, or Bows. (300)" 
	"reward" "•Double XP" 
}}
{{ $challenge_3 := sdict 
	"title" "Lenses in Focus" 
	"description" "Go to the Prismatic Recaster and discover how to unlock (10) more lenses." 
	"reward" "•XP" 
}}
{{ $challenge_4 := sdict 
	"title" "Clearing the Inner Circle" 
	"description" "Defeat (40) Primeval Envoys in Gambit." 
	"reward" "•4x XP  •Bright Dust" 
}}
{{ $challenge_5 := sdict 
	"title" "Trial by Firing Squad" 
	"description" "Win multiple (7) rounds in the Trials of Osiris." 
	"reward" "•4x XP  •Bright Dust  •Trials of Osiris Weapon" 
}}
{{ $challenge_6 := sdict 
	"title" "Decisive Strike" 
	"description" "Acquire the Panacea ornament for the Salvager's Salvo Grenade Launcher." 
	"reward" "•Double XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_8" (sdict 
	"title" "Seasonal Challenges - Week 8" 
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
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_8").Value)}}