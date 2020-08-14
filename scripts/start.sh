#!/bin/bash

set -ex

PROJECT_ID=webserver1-283520
BINARY_NAME=apiserver
SERVICE_NAME=apiserver-service

#docker build -t gcr.io/${PROJECT_ID}/${BINARY_NAME}:v1 -f ../Dockerfile-apiserver ..
docker images

yes | gcloud auth configure-docker
docker push gcr.io/${PROJECT_ID}/${BINARY_NAME}:v1

gcloud config set project $PROJECT_ID
gcloud config set compute/zone us-central1-a
#gcloud container clusters create my-cluster1
gcloud compute instances list

#kubectl create deployment ${BINARY_NAME} --image=gcr.io/${PROJECT_ID}/${BINARY_NAME}:v1
kubectl scale deployment ${BINARY_NAME} --replicas=3
#kubectl autoscale deployment ${BINARY_NAME} --cpu-percent=80 --min=1 --max=5
kubectl get pods

kubectl expose deployment ${BINARY_NAME} --name=${SERVICE_NAME} --type=LoadBalancer --port 80 --target-port 8080
kubectl get service
