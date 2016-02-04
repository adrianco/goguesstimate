curl -vs http://guesstimate.herokuapp.com/spaces -X POST --data @./$1.guess -H "Content-Type: application/json" -H "Authorization: Bearer $GUESS"
