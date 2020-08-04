all:
	docker build . -t registry.heroku.com/life-spacy-server/web
	docker push registry.heroku.com/life-spacy-server/web
	heroku container:release -a life-spacy-server web
test:
	@curl https://life-spacy-server.herokuapp.com/ner -H 'Content-Type: application/json' -d '{"sections": ["My name is John Doe. I grew up in California."]}'
