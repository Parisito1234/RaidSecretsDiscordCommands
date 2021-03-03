{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/DestinyActivityModeDefinition_e35792b49b249ca5dcdb1e7657ca42b6.png" }}
{{$embed := cembed 
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "Trials of Osiris loot table" "value" "[Loot Rotation Spreadsheet](https://docs.google.com/spreadsheets/u/1/d/e/2PACX-1vTtZbHHw3ncerNkBsOL5ckJf1Fi_DrVOyl7maIAIIpPj-FSmFWoZZDhcva3xkLz_ExQ70XpWPe4c6BH/pubhtml?gid=0&single=true)" "inline" false) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
}}
{{ sendMessage nil $embed }}