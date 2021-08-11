{{ /* Trigger: Command: git */ }}

{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/8ff7269628a34bf7adc71ee0b518f9a2.png" }}

{{$embed := cembed 
    "color" 1772743
    "fields" (cslice 
        (sdict "name" " == __Documentation__ == " "value" "[Monteven's Docs](https://github.com/MontagueM/DestinyDatamining)" "inline" false) 
	(sdict "name" "Model Extractor" "value" "[Monteven's Dynamic Extractor](https://github.com/MontagueM/MontevenDynamicExtractor)" "inline" true) 
        (sdict "name" "Texture Tools" "value" "[Monteven's Destiny Tex Converter](https://github.com/MontagueM/DestinyTexConv)" "inline" true) 
        (sdict "name" " == __API:__ == " "value" "[Official Bungie API](https://bungie-net.github.io/multi/)" "inline" false) 
        (sdict "name" "API - JS" "value" "[Parisito's Google App Script Library](https://github.com/Parisito1234/AppScriptDestinyAPI)" "inline" true) 
        (sdict "name" "API - Python" "value" "[Tadpolefeet's Python Library](https://github.com/TadpolefeetDestiny2/PythonDestinyAPI) \n[Nxtlo's Python Library](https://github.com/nxtlo/aiobungie)" "inline" true) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
}}

{{ sendMessage nil $embed }}