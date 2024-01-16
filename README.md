# go-todo

- ビルド

  ```
  docker build -t go-todo .
  ```

- コンテナ起動

  ```
  docker-compose up -d
  ```

- kubernetes 起動

  ```
  kubectl apply -f configmap.yaml
  kubectl apply -f storage-class-ssd.yaml
  kubectl apply -f persistentvolume.yaml
  kubectl apply -f db.yaml
  kubectl apply -f api.yaml
  ```
