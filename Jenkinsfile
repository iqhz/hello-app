pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "iqhz/hello-app"
        REGISTRY_CREDENTIALS = 'dockerhub-credentials'      // Username & password Docker Hub
        KUBECONFIG_CREDENTIALS = 'kubeconfig-rke2'          // Kubeconfig file (disimpan di Jenkins)
    }

    stages {
        stage('Checkout') {
            steps {
                git 'git@github.com:iqhz/hello-app.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t $DOCKER_IMAGE:latest .'
                }
            }
        }

        stage('Login to Docker Hub') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: env.REGISTRY_CREDENTIALS,
                    usernameVariable: 'DOCKER_USERNAME',
                    passwordVariable: 'DOCKER_PASSWORD'
                )]) {
                    sh 'echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin'
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                sh 'docker push $DOCKER_IMAGE:latest'
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                withCredentials([file(credentialsId: env.KUBECONFIG_CREDENTIALS, variable: 'KUBECONFIG')]) {
                    sh 'kubectl --kubeconfig=$KUBECONFIG apply -f k8s/'
                }
            }
        }
    }

    post {
        success {
            echo '✅ Build, push, and deploy completed successfully!'
        }
        failure {
            echo '❌ Build failed. Please check the logs.'
        }
    }
}
