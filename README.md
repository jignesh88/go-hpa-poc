# Kubernetes Horizontal Pod Autoscaling (HPA) Demo

This repository contains a demonstration of Kubernetes Horizontal Pod Autoscaling (HPA) using a simple Go application that can generate CPU load on demand.

## Components

- `main.go`: A Go web server with endpoints for health checks and CPU load generation
- `Dockerfile`: Multi-stage build for creating a small, efficient container image
- `deployment.yaml`: Kubernetes deployment and service definitions with resource limits and probes
- `hpa.yaml`: HPA configuration with CPU-based scaling and customized behavior
- `demo-scripts.sh`: Helper scripts for running and demonstrating the HPA functionality

## Usage

1. Build the Docker image:
   ```bash
   docker build -t hpa-demo-app:latest .
   ```

2. If using minikube, load the image:
   ```bash
   minikube image load hpa-demo-app:latest
   ```

3. Deploy the application:
   ```bash
   kubectl apply -f deployment.yaml
   kubectl apply -f hpa.yaml
   ```

4. Ensure metrics-server is installed (if not already):
   ```bash
   kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
   ```

5. Access the application and monitor HPA:
   ```bash
   # Terminal 1: Port forward to access the service
   kubectl port-forward svc/hpa-demo-service 8080:80
   
   # Terminal 2: Watch the HPA
   kubectl get hpa hpa-demo-app-hpa --watch
   
   # Terminal 3: Generate load
   curl "http://localhost:8080/load?duration=180&cores=1"
   ```

## Application Endpoints

- `/`: Basic information about the application
- `/load?duration=X&cores=Y`: Generate CPU load for X seconds using Y cores
- `/health`: Health check endpoint
- `/ready`: Readiness check endpoint

## HPA Configuration

The HPA is configured to:
- Maintain between 2 and 10 replicas
- Target 80% CPU utilization
- Scale up quickly (0-second stabilization window)
- Scale down conservatively (5-minute stabilization window)

See `hpa.yaml` for complete configuration details.

## License

MIT