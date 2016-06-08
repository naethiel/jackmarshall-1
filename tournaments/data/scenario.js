db = db.getSiblingDB('jackmarshall')
db.scenario.drop()
db.scenario.insert({
		"name": "Destruction",
		"year": 2015,
		"link": "http://google.com"
	}
)
db.scenario.insert({
		"name": "Kill on sight",
		"year": 2015,
		"link": "http://google.com"
	}
)
