{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/DestinyActivityModeDefinition_e35792b49b249ca5dcdb1e7657ca42b6.png" }}
{{$embed := cembed 
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "Trials of Osiris loot table" "value" "[Loot Rotation Spreadsheet](https://docs.google.com/spreadsheets/d/1cbIuRPujrM4R-Q-C9GcRhdFfkVFut8k8bNkMB87FujA/edit?usp=sharing)" "inline" false) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
}}
{{ sendMessage nil $embed }}