{{ /* Trigger: Exact Match: ~dbtest */ }}
{{ /* Currently deletes first 10 pattern matches of "embed_" - underscore is wildcard due to PostgreSQL - this is a testing command. Treat it as such. */ }}

{{ dbDelMultiple (sdict "pattern" "embed_") 10 0 }}