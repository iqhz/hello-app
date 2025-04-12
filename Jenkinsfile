pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "iqhz/hello-app"
        REGISTRY_CREDENTIALS = 'dockerhub-credentials'   // Ganti dengan ID credentials Docker Hub kamu
        KUBECONFIG_CREDENTIALS = 'kubeconfig-rke2'       // ID credentials untuk kubeconfig
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
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
                withCredentials([file(credentialsId: env.KUBECONFIG_CREDENTIALS, variable: 'KUBECONFIG')]) {
                    // Debug kubectl config path
                    echo "Kubeconfig path: $KUBECONFIG"

                    // Make sure the file is correctly handled by Jenkins
                    sh "cat $KUBECONFIG"  // Print kubeconfig for verification
                    
                    // Use kubectl with correct kubeconfig path
                    sh 'kubectl --kubeconfig=$KUBECONFIG version'
                    sh 'kubectl --kubeconfig=$KUBECONFIG config get-contexts'

                    // Apply all YAML files in k8s directory
                    dir('k8s') {
                        sh 'kubectl --kubeconfig=$KUBECONFIG apply -f .'
                    }
                }
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
