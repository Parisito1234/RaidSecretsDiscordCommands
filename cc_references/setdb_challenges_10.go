{{ $challenge_1 := sdict 
	"title" "…The Harder They Fall" 
	"description" "Defeat (20) Elite or Boss Cabal in the Proving Grounds strike." 
	"reward" "•XP" 
}}
{{ $challenge_2 := sdict 
	"title" "Explosive Conclusions" 
	"description" "Use a Rocket Launcher or grenades to defeat combatants anywhere in the system. Bonus progress for rapidly defeating them and for defeating Cabal. (600)" 
	"reward" "•Double XP" 
}}
{{ $challenge_3 := sdict 
	"title" "Lenses in Focus" 
	"description" "Go to the Prismatic Recaster and discover how to unlock (15) more lenses." 
	"reward" "•XP" 
}}
{{ $challenge_4 := sdict 
	"title" "In Your Element " 
	"description" "Defeat (80) Guardians in the Iron Banner with Elemental takedowns. Earn bonus progress for Stasis takedowns." 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_5 := sdict 
	"title" "Warrior from Beyond" 
	"description" "In strikes, defeat combatants with elemental final blows. Earn bonus progress by defeating combatants with Stasis. (300)" 
	"reward" "•Double XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_10" (sdict 
	"title" "Seasonal Challenges - Week 10" 
	"color" 1772743
	"fields" (cslice 
		(sdict "name" ($challenge_1.Get "title") "value" (joinStr "" ($challenge_1.Get "description") " **|** `" ($challenge_1.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_2.Get "title") "value" (joinStr "" ($challenge_2.Get "description") " **|** `" ($challenge_2.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_3.Get "title") "value" (joinStr "" ($challenge_3.Get "description") " **|** `" ($challenge_3.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_4.Get "title") "value" (joinStr "" ($challenge_4.Get "description") " **|** `" ($challenge_4.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_5.Get "title") "value" (joinStr "" ($challenge_5.Get "description") " **|** `" ($challenge_5.Get "reward") "`") "inline" false) 
	) 
)}}
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_10").Value)}}