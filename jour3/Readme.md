# LECTURES
curl -s -w '%{http_code}\n' localhost:8080/stations                       # 200
curl -s -w '%{http_code}\n' localhost:8080/stations/FR-BOR-001            # 200
curl -s -w '%{http_code}\n' localhost:8080/stations/FR-BOR-001/observations  # 200
curl -s -w '%{http_code}\n' localhost:8080/stations/XX-NOPE                # 404

# ÉCRITURES
curl -s -w '%{http_code}\n' -X POST localhost:8080/stations \
-H 'Content-Type: application/json' \
-d '{"id":"T-001","name":"Test"}'                                  # 201

curl -s -w '%{http_code}\n' -X POST localhost:8080/stations \
-d '{"id":"T-001","name":"X"}' -H 'Content-Type: application/json'   # 409

curl -s -w '%{http_code}\n' -X PUT localhost:8080/stations/T-001 \
-d '{"name":"T2"}' -H 'Content-Type: application/json'                 # 200

curl -s -w '%{http_code}\n' -X DELETE localhost:8080/stations/T-001        # 204
curl -s -w '%{http_code}\n' -X DELETE localhost:8080/stations/T-001        # 404
