#!/bin/bash

# kubectl apply -f "https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml"
# kubectl apply -f https://raw.githubusercontent.com/rabbitmq/cluster-operator/main/docs/examples/hello-world/rabbitmq.yaml

# # kubectl get rabbitmqclusters.rabbitmq.com

# username="$(kubectl get secret hello-world-default-user -o jsonpath='{.data.username}' | base64 --decode)"
# echo "username: $username"
# password="$(kubectl get secret hello-world-default-user -o jsonpath='{.data.password}' | base64 --decode)"
# echo "password: $password"

# kubectl port-forward "service/hello-world" 15672


username="$(kubectl get secret hello-world-default-user -o jsonpath='{.data.username}' | base64 --decode)"
password="$(kubectl get secret hello-world-default-user -o jsonpath='{.data.password}' | base64 --decode)"
service="$(kubectl get service hello-world -o jsonpath='{.spec.clusterIP}')"
kubectl run perf-test --image=pivotalrabbitmq/perf-test -- --uri amqp://$username:$password@$service