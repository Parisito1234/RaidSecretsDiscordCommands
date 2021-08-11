{{ /* Trigger: Command: api */ }}

{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/8ff7269628a34bf7adc71ee0b518f9a2.png" }}

{{$embed := cembed 
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "Destiny API tools" "value" "[data.destinysets.com](https://data.destinysets.com/) \n[Light.gg](https://www.light.gg/) \n[Destiny Report Archive](https://archive.destiny.report/)" "inline" true) 
        (sdict "name" "Documentation" "value" "[Bungie API Docs](https://bungie-net.github.io/) \n[Bungie Git Page](https://github.com/Bungie-net/api)" "inline" true) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
}}

{{ sendMessage nil $embed }}