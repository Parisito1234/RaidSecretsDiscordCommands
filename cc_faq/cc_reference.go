{{$embed := cembed 
    "title" "Reference commands in this server:"
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "-lore" "value" "Instant search links to Ishtar Collective\nUsage: `-lore [Opt: 'categories|books|items'] (search query)`" "inline" false) 
        (sdict "name" "-challenges" "value" "Provides a link to current seasonal challenges. \nUsage: `-challenges`" "inline" false) 
        (sdict "name" "-twab" "value" "Link to TWAB webpage" "inline" true) 
        (sdict "name" "-patchnotes" "value" "Link to Bungie Updates webpage" "inline" true) 
        (sdict "name" "-roadmap" "value" "Current Seasonal Roadmap" "inline" true) 
        (sdict "name" "-trials" "value" "Trials of Osiris loot table link" "inline" true) 
        (sdict "name" "-trailers" "value" "Currently relevant trailers for Destiny 2" "inline" true) 
        (sdict "name" "-presskit" "value" "Destiny 2 Press Kit Dropbox link" "inline" true) 
        (sdict "name" "-wish15" "value" "Wish 15 status and links" "inline" true) 
    ) 
}}
 
{{ sendMessage nil $embed }}