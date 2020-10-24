main.go contains the logic to compute number of nodes running in the current cluster

I have not included vendor folder that contains packages due to size constraints, 
they should contain following packages
github.com 
golang.org 
google.golang.org
gopkg.in
k8s.io
sigs.k8s.io

Command to build docker image locally

```
  docker build -t <image_name> .
```

To run previously built docker image use, presuming your current context is minikube, attaching config and certificates as volumes

```
docker run -it -p 8080:8080 -v /$HOME/.kube:/root/.kube -v /$HOME/.minikube:/$HOME/.minikube --name <image_name> <image_name>
```

Once you have the container running, ssh into the container using following commands

```
 docker exec -ti <container_id> sh
```

Then use below command to query number of pods currently present in the current cluster
```
    wget http://localhost:8081/json
```
above command returns a file called json, contents of this file contains number of pods currently available


For additional testing, create a dummy pod using k8.yaml that's included in the repo

```
k apply -f k8.yaml -n<namespace>
```

above command creates additional pod, use http://localhost:8081/json to get updated result
