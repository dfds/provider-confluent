#!/bin/bash

# Go to project root
pushd $(git rev-parse --show-toplevel)

# Start minikube cluster
minikube start --cpus 2 --memory 2048

# Create crossplane-system namespace
kubectl create namespace crossplane-system

# Add and install crossplane helm chart into crossplane-system namespace
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update
helm install crossplane --namespace crossplane-system crossplane-stable/crossplane

# Setup build dependencies
make submodules

# Start dev
make dev

# Install provider config
envsubst < examples/provider/config.yaml | kubectl apply -f -

# Go back
popd