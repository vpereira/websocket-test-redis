Test to update table cel based on the results coming from websocket backed by
redis

### To run
```docker-compose up```

- Create some jobs:

```curl -X POST http://localhost:3000/create.json```

- Delete all jobs:

```curl -X DELETE http://localhost:3000/delete.json```

- Open Browser on http://localhost:3000
- curl a json to the update json, i.e:

```curl -H "Accept: application/json" -X POST -d '{"id": "5ed2b107-c3f9-41d5-860e-9e078a3cf2a7", "status": "failed"}' localhost:3000/update.json```

- Watch it update in the webui
