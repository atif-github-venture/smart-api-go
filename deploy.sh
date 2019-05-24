#!/usr/bin/env bash
#to run this ./deploy.sh
#if having issue with permission chmod +x ./deploy.sh

echo y | docker system prune

cd action-list
docker build -t action-list .
docker tag action-list atifdockerventure/action-list
docker push atifdockerventure/action-list:latest
cd ..

cd create-project
docker build -t create-project .
docker tag create-project atifdockerventure/create-project
docker push atifdockerventure/create-project:latest
cd ..

cd identifier-repository
docker build -t identifier-repository .
docker tag identifier-repository atifdockerventure/identifier-repository
docker push atifdockerventure/identifier-repository:latest
cd ..

cd test-repository
docker build -t test-repository .
docker tag test-repository atifdockerventure/test-repository
docker push atifdockerventure/test-repository:latest
cd ..

cd test-config
docker build -t test-config .
docker tag test-config atifdockerventure/test-config
docker push atifdockerventure/test-config:latest
cd ..

cd data-setup
docker build -t data-setup .
docker tag data-setup atifdockerventure/data-setup
docker push atifdockerventure/data-setup:latest
cd ..

cd kube-cluster
kubectl delete service/action-list
kubectl delete service/create-project
kubectl delete service/identifier-repository
kubectl delete service/test-repository
kubectl delete service/test-config
kubectl delete service/data-setup

kubectl delete deployment.apps/action-list
kubectl delete deployment.apps/create-project
kubectl delete deployment.apps/identifier-repository
kubectl delete deployment.apps/test-repository
kubectl delete deployment.apps/test-config
kubectl delete deployment.apps/data-setup

kubectl apply -f mongo-service.yml
cd api
kubectl create -f action-list-deploy-svc.yml
kubectl create -f create-project-deploy-svc.yml
kubectl create -f identifier-repository-deploy-svc.yml
kubectl create -f test-repository-deploy-svc.yml
kubectl create -f test-config-deploy-svc.yml
kubectl create -f data-setup-deploy-svc.yml
cd ..
cd loadbalancer
kubectl apply -f nginx.yml
kubectl apply -f ingress.yml

minikube ip
kubectl get all
kubectl describe ing


#Note:
#Have the mongo db deployed and have the nginx.yml applied
#kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
#run -> minikube addons enable ingress
#run -> curl -kL http://minikube_ip/service_name
#source: https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ingress-guide-nginx-example.html
#for public ip example -> https://dzone.com/articles/ingress-controllers-for-kubernetes