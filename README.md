# Backend

## Commands of application
Use this command in the directory
- `docker-compose up -d --build` - to run an application
- `docker-compose down` - to stop an application


# _In Postman use_
to create people http://localhost:80/api/people
```
{
    "name" : "Dima",
    "phone_number" : "38098088829590",
    "birth_year" : "2003"
}
```
## Comands
- http://localhost:80/api/people - create people
- http://localhost:80/api/people/_id - get people
- http://localhost:80/api/peoples - get all peoples
- http://localhost:80/api/people/_id - delete people
- http://localhost:80/api/deleteall - delete all peoples
- http://localhost:80/api/people/_id - update people


