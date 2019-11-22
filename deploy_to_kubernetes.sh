#!/usr/bin/env bash
#to run this ./deploy.sh
#if having issue with permission chmod +x ./deploy.sh

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

kubectl get all
kubectl describe ing