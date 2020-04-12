
#http POST :8080/users id=1 first_name=Ankit last_name=Agrawal email=aankittcoolest

curl -X POST localhost:8080/users -d '{"id": 123, "first_name": "Ankit", "last_name": "Agrawal"}'
#curl -X POST localhost:8080/users -d '{"id": 123}'
