{{$embed := cembed 
    "title" "Datamining Discussion Commands:"
    "description" "This channel has specific overrides for `-reference` and any commands that have the same trigger as the rest of the server."
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "-git" "value" "Useful code repositories for datamining tools and documentation" "inline" false) 
        (sdict "name" "-api" "value" "API documentation and 3rd party tools" "inline" false) 
    ) 
}}
 
{{ sendMessage nil $embed }}