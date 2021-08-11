{{ /* Trigger: Command: challenges */ }}

{{ $avatar := (dbGet 0 "currentSeasonIcon").Value }}
{{$embed := cembed 
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "All weekly seasonal triumphs" "value" "[Vault of Glass](https://vaultof.glass/weeklies)\n*Courtesy of nev*" "inline" false) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
}}
{{ sendMessage nil $embed }}
