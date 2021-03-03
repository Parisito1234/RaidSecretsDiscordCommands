{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/8ff7269628a34bf7adc71ee0b518f9a2.png" }}
 
{{$embed := cembed 
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "Documentation" "value" "[Monteven's Docs](https://github.com/MontagueM/DestinyDatamining)" "inline" false) 
        (sdict "name" "Text Tools" "value" "[Monteven's Destiny Text Converter](https://github.com/MontagueM/DestinyTexConv)" "inline" false) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
}}
 
{{ sendMessage nil $embed }}