#!/usr/bin/env bash
#to run this ./deploy.sh
#if having issue with permission chmod +x ./deploy.sh

#echo y | docker system prune
#
#cd agere-lystan
#docker build -t agere-lystan .
#docker tag agere-lystan atifdockerventure/agere-lystan
#docker push atifdockerventure/agere-lystan:latest
#cd ..
#
#cd creare-project
#docker build -t creare-project .
#docker tag creare-project atifdockerventure/creare-project
#docker push atifdockerventure/creare-project:latest
#cd ..
#
#cd detail-testa
#docker build -t detail-testa .
#docker tag detail-testa atifdockerventure/detail-testa
#docker push atifdockerventure/detail-testa:latest
#cd ..
#
#cd identifier-reponere
#docker build -t identifier-reponere .
#docker tag identifier-reponere atifdockerventure/identifier-reponere
#docker push atifdockerventure/identifier-reponere:latest
#cd ..
#
#cd makros-testa
#docker build -t makros-testa .
#docker tag makros-testa atifdockerventure/makros-testa
#docker push atifdockerventure/makros-testa:latest
#cd ..
#
#cd mikros-testa
#docker build -t mikros-testa .
#docker tag mikros-testa atifdockerventure/mikros-testa
#docker push atifdockerventure/mikros-testa:latest
#cd ..

cd kube-cluster
kubectl delete service/mongo
kubectl delete service/agere-lystan
kubectl delete service/creare-project
kubectl delete service/detail-testa
kubectl delete service/identifier-reponere
kubectl delete service/makros-testa
kubectl delete service/mikros-testa

kubectl delete deployment.apps/agere-lystan
kubectl delete deployment.apps/creare-project
kubectl delete deployment.apps/detail-testa
kubectl delete deployment.apps/identifier-reponere
kubectl delete deployment.apps/makros-testa
kubectl delete deployment.apps/mikros-testa

kubectl create -f mongo-service.yml
kubectl create -f agere-lystan-deploy-svc.yml
kubectl create -f creare-project-deploy-svc.yml
kubectl create -f detail-testa-deploy-svc.yml
kubectl create -f identifier-reponere-deploy-svc.yml
kubectl create -f makros-testa-deploy-svc.yml
kubectl create -f mikros-testa-deploy-svc.yml
kubectl apply -f ingress.yml

minikube ip
kubectl get all
kubectl describe ing