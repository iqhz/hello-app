pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "iqhz/hello-app"
        REGISTRY_CREDENTIALS = 'dockerhub-credentials'   // Jenkins credentials ID
        KUBECONFIG_CREDENTIALS = 'kubeconfig-rke2'       // Opsional kalau pakai credentials
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

        stage('Push Image') {
            steps {
                sh 'docker push $DOCKER_IMAGE:latest'
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                // Jika kubeconfig disimpan di Jenkins credentials:
                // withCredentials([file(credentialsId: env.KUBECONFIG_CREDENTIALS, variable: 'KUBECONFIG')]) {
                //     sh 'kubectl --kubeconfig=$KUBECONFIG apply -f k8s/'
                // }

                // Kalau kubeconfig sudah tersedia di ~/.kube/config:
                sh 'kubectl apply -f k8s/'
            }
        }
    }

    post {
        success {
            echo '✅ Deployment success!'
        }
        failure {
            echo '❌ Something went wrong...'
        }
    }
}
