# GoLang-Redis-Cluster

**Run on docker the redis services**
```bash
  docker-compose up -d redis1 redis2 redis3 redis4 redis5 redis6
```

**Find ip addresses of redis containers as below command**
```bash
  docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis1 // => 172.23.0.2
```

**Enter to the container with exec command**
```bash
  docker exec -it redis1 bash
```

**Create Redis Cluster**
```bash
  redis-cli --cluster create 172.23.0.2:6379 172.23.0.3:6379 172.23.0.4:6379 172.23.0.5:6379 172.23.0.6:6379 172.23.0.7:6379 --cluster-replicas 1
```

**Docker build GoLang App**
```bash
  docker build -t redis-cluster-app .
```

**Run docker-compose command**
```bash
  docker-compose up -d
```
