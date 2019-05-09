#!/usr/bin/env bash
#to run this ./deploy.sh
#if having issue with permission chmod +x ./deploy.sh

echo y | docker system prune

cd agere-lystan
docker build -t agere-lystan .
docker tag agere-lystan atifdockerventure/agere-lystan
docker push atifdockerventure/agere-lystan:latest
cd ..

cd creare-project
docker build -t creare-project .
docker tag creare-project atifdockerventure/creare-project
docker push atifdockerventure/creare-project:latest
cd ..

cd identifier-reponere
docker build -t identifier-reponere .
docker tag identifier-reponere atifdockerventure/identifier-reponere
docker push atifdockerventure/identifier-reponere:latest
cd ..

cd makros-testa
docker build -t makros-testa .
docker tag makros-testa atifdockerventure/makros-testa
docker push atifdockerventure/makros-testa:latest
cd ..

cd mikros-testa
docker build -t mikros-testa .
docker tag mikros-testa atifdockerventure/mikros-testa
docker push atifdockerventure/mikros-testa:latest
cd ..

cd testa-configurare
docker build -t testa-configurare .
docker tag testa-configurare atifdockerventure/testa-configurare
docker push atifdockerventure/testa-configurare:latest
cd ..

cd data-setup
docker build -t data-setup .
docker tag data-setup atifdockerventure/data-setup
docker push atifdockerventure/data-setup:latest
cd ..

cd kube-cluster
kubectl delete service/agere-lystan
kubectl delete service/creare-project
kubectl delete service/identifier-reponere
kubectl delete service/makros-testa
kubectl delete service/mikros-testa
kubectl delete service/testa-configurare
kubectl delete service/data-setup

kubectl delete deployment.apps/agere-lystan
kubectl delete deployment.apps/creare-project
kubectl delete deployment.apps/identifier-reponere
kubectl delete deployment.apps/makros-testa
kubectl delete deployment.apps/mikros-testa
kubectl delete deployment.apps/testa-configurare
kubectl delete deployment.apps/data-setup

kubectl apply -f mongo-service.yml
cd api
kubectl create -f agere-lystan-deploy-svc.yml
kubectl create -f creare-project-deploy-svc.yml
kubectl create -f identifier-reponere-deploy-svc.yml
kubectl create -f makros-testa-deploy-svc.yml
kubectl create -f mikros-testa-deploy-svc.yml
kubectl create -f testa-configurare-deploy-svc.yml
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