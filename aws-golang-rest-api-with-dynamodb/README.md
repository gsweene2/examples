# aws-golang-rest-api-with-dynamodb

Build & Deploy
```
make deploy
```

# CRUD Operations

## Create

```
curl --request POST \
  --url https://fz3n8nstdf.execute-api.us-east-1.amazonaws.com/dev/todos \
  --header 'Content-Type: application/json' \
  --data '{
  "Title": "Walk the Dog",
  "Details": "Complete before 11am"
}'

curl --request POST \
  --url https://fz3n8nstdf.execute-api.us-east-1.amazonaws.com/dev/todos \
  --header 'Content-Type: application/json' \
  --data '{
  "Title": "Mow the Lawn",
  "Details": "Remember to buy gas"
}'
```

## Read

```
curl --request GET \
  --url https://fz3n8nstdf.execute-api.us-east-1.amazonaws.com/dev/todos/{id}
```

## Update

```
curl --request PUT \
  --url https://fz3n8nstdf.execute-api.us-east-1.amazonaws.com/dev/todos/0d2263b7-c62d-4df6-8503-bb16ee8dd81 \
  --header 'Content-Type: application/json' \
  --data '{
  "title": "Updated title",
  "details": "Updated details"
}'
```

## List

```
curl --request GET \
  --url https://fz3n8nstdf.execute-api.us-east-1.amazonaws.com/dev/todos
```


## Delete

```
curl --request DELETE \
  --url https://fz3n8nstdf.execute-api.us-east-1.amazonaws.com/dev/todos/0d2263b7-c62d-4df6-8503-bb16ee8dd81
```
