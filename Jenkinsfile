pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "iqhz/hello-app"
        REGISTRY_CREDENTIALS = 'dockerhub-credentials'
        KUBECONFIG_CREDENTIALS = 'kubeconfig-rke2'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
                script {
                    // Ambil commit ID dan set full image tag
                    def commitId = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
                    env.DOCKER_IMAGE_TAG = "${DOCKER_IMAGE}:${commitId}"
                    echo "Docker Image Tag: ${env.DOCKER_IMAGE_TAG}"
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh "docker build -t ${env.DOCKER_IMAGE_TAG} ."
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
                sh "docker push ${env.DOCKER_IMAGE_TAG}"
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                withCredentials([file(credentialsId: env.KUBECONFIG_CREDENTIALS, variable: 'KUBECONFIG')]) {
                    // Debug kubeconfig
                    echo "Kubeconfig path: $KUBECONFIG"
                    sh "kubectl --kubeconfig=$KUBECONFIG version"
                    sh "kubectl --kubeconfig=$KUBECONFIG config get-contexts"

                    // Update deployment image
                    sh "kubectl --kubeconfig=$KUBECONFIG set image deployment/hello-app hello-app=${env.DOCKER_IMAGE_TAG} --record"
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
