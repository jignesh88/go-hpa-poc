# Build the Docker image
docker build -t hpa-demo-app:latest .

# Load the image into minikube (if you're using minikube)
# minikube image load hpa-demo-app:latest

# Apply the deployment
kubectl apply -f deployment.yaml

# Apply the HPA
kubectl apply -f hpa.yaml

# Check that everything is running
kubectl get pods
kubectl get hpa

# Ensure metrics-server is installed (if not already installed)
# kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

# Port-forward to access the service locally
kubectl port-forward svc/hpa-demo-service 8080:80

# In a separate terminal, watch the HPA
kubectl get hpa hpa-demo-app-hpa --watch

# In yet another terminal, generate load to trigger scaling
# (Adjust the URL if you're accessing it differently)
curl "http://localhost:8080/load?duration=180&cores=1"

# Wait and observe the HPA scaling up the deployment
# After a few minutes, the load will stop and you should see it scale back down