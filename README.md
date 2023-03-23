# Backend

## To start server
- docker build -t cz-api:v0.1 .
- docker run --name cz-api -p 80:80 cz-api:v0.1

# _In Postman use_
to create people http://localhost:8080/api/people
```
{
    "name": "dima",
    "phone_number": phone_number,
    "birth_year": 
}
```
## Comands
- http://localhost:8080/api/people - create people
- http://localhost:8080/api/people/_id - get people
- http://localhost:8080/api/peoples - get all peoples
- http://localhost:8080/api/people/_id - delete people
- http://localhost:8080/api/deleteall - delete all peoples
- http://localhost:8080/api/people/_id - update people


