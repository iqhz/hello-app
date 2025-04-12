pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "iqhz/hello-app"
        REGISTRY_CREDENTIALS = 'dockerhub-credentials'   // Ganti dengan ID credentials Docker Hub kamu
        KUBECONFIG_CREDENTIALS = 'kubeconfig-rke2'       // ID credentials untuk kubeconfig
        COMMIT_ID = ""                                    // Variabel untuk menyimpan commit ID
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
                script {
                    // Ambil commit ID dari Git dan simpan sebagai variabel
                    env.COMMIT_ID = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
                    echo "Commit ID: ${env.COMMIT_ID}"
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Build Docker image dengan commit ID sebagai tag
                    sh "docker build -t $DOCKER_IMAGE:${env.COMMIT_ID} ."
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
                // Push Docker image dengan commit ID
                sh "docker push $DOCKER_IMAGE:${env.COMMIT_ID}"
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
                        // Update deployment with the new image
                        sh "kubectl --kubeconfig=$KUBECONFIG set image deployment/hello-app hello-app=$DOCKER_IMAGE:${env.COMMIT_ID} --record"
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
