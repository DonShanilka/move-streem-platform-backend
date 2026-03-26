#!/bin/bash

# Configuration
SERVICES=("movie-service" "auth-service" "user-service" "admin-service" "genres-service" "payment-service" "tvSeries-service")
DOCKER_USER="" # Leave empty for local use

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
BACKEND_DIR="$( cd "$SCRIPT_DIR/.." && pwd )"

echo "🚀 Starting builds in $BACKEND_DIR..."

for SERVICE in "${SERVICES[@]}"; do
    echo "---------------------------------------------------"
    echo "📦 Building $SERVICE..."
    
    # Adjust directory name for tvSeries-service if needed (manifest uses tvseries-service, folder is tvSeries-service)
    DIR_NAME=$SERVICE
    IMAGE_NAME=$(echo $SERVICE | tr '[:upper:]' '[:lower:]')
    
    cd "$BACKEND_DIR/services/$DIR_NAME"
    
    docker build -t "$IMAGE_NAME:latest" .
    
    if [ $? -ne 0 ]; then
        echo "❌ Failed to build $SERVICE"
        exit 1
    fi
    
    echo "✅ Built $IMAGE_NAME:latest"
    
    # If using minikube, load the image
    if command -v minikube &> /dev/null; then
        echo "📥 Loading $IMAGE_NAME:latest into minikube..."
        minikube image load "$IMAGE_NAME:latest"
    fi

    # If using kind, load the image
    # if command -v kind &> /dev/null; then
    #     echo "📥 Loading $IMAGE_NAME:latest into kind..."
    #     kind load docker-image "$IMAGE_NAME:latest"
    # fi
done

echo "---------------------------------------------------"
echo "✨ All images built and loaded!"
echo "🔄 Restarting pods..."
kubectl rollout restart deployment -n default
