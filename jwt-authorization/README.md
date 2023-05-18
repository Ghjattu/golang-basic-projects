# JWT Authorization

## Simple REST API Design
In this project, we create 2 **public** endpoints:
1. `/api/register`
2. `/api/login` : this route will be used to authorize the user by providing `username` and `password`, then generate and return a JSON Web Token.

and 1 **protected** endpoint that will be protected by JWT:
1. `/api/users` : this route returns information for all users.

## Start
Create the `.env` file in the root directory, add:
```
SECRET_KEY=<your-secret-key>
```
SECRET_KEY is your secret string for signing the token.

Then just `go run main.go` and test these APIs using an API client.