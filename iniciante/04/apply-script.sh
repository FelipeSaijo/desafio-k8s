#!/bin/bash

TEMPLATES='https://raw.githubusercontent.com/FelipeSaijo/desafio-k8s/refs/heads/main/iniciante/04-network-issue/deployment.yaml','https://raw.githubusercontent.com/FelipeSaijo/desafio-k8s/refs/heads/main/iniciante/04-network-issue/service.yaml','https://raw.githubusercontent.com/FelipeSaijo/desafio-k8s/refs/heads/main/iniciante/04-network-issue/netwokpolicies.yaml'

echo "Applying Templates..."

for template in $TEMPLATES
do
    kubectl apply -f $template > /dev/null 2>&1
    
    if [ $? -ne 0 ]; then
        echo "Error applying $template. Exiting..."
        exit 1
    fi
done

echo "All templates applied successfully!"
