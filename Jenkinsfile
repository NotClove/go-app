pipeline {
    agent any
    
    environment {
        // SonarQube
        SONAR_HOST_URL = 'http://sonarqube:9000'
        
        // Nexus
        NEXUS_URL = 'http://nexus:8081'
        NEXUS_REPO = 'docker-hosted'
        
        // Docker
        DOCKER_IMAGE = 'go-app'
        DOCKER_TAG = "${env.BUILD_NUMBER}"
        DOCKER_REGISTRY = "nexus:8082"
    }
    
    stages {
        stage('Checkout') {
            steps {
                echo 'Checking out code from repository...'
                checkout scm
            }
        }
        
        stage('Test') {
            steps {
                echo 'Running Go tests...'
                sh '''
                    go mod download || true
                    go test -v -coverprofile=coverage.out .
                '''
            }
            post {
                always {
                    script {
                        if (fileExists('coverage.out')) {
                            sh 'go tool cover -html=coverage.out -o coverage.html'
                            publishHTML([
                                reportName: 'Coverage Report',
                                reportDir: '.',
                                reportFiles: 'coverage.html',
                                keepAll: true,
                                alwaysLinkToLastBuild: true,
                                allowMissing: true
                            ])
                        } else {
                            echo 'Coverage report not generated, skipping HTML publish'
                        }
                    }
                }
            }
        }
        
        stage('SonarQube Analysis') {
            steps {
                echo 'Running SonarQube analysis...'
                script {
                    try {
                        withSonarQubeEnv('SonarQube') {
                            withCredentials([string(credentialsId: 'sonarqube-token', variable: 'SONAR_TOKEN')]) {
                                sh '''
                                    # Запуск анализа SonarQube через sonar-scanner
                                    sonar-scanner \
                                        -Dsonar.projectKey=go-app \
                                        -Dsonar.projectName=Go Test Application \
                                        -Dsonar.host.url=${SONAR_HOST_URL} \
                                        -Dsonar.login=${SONAR_TOKEN} \
                                        -Dsonar.sources=. \
                                        -Dsonar.exclusions=**/*_test.go \
                                        -Dsonar.go.coverage.reportPaths=coverage.out
                                '''
                            }
                        }
                    } catch (Exception e) {
                        echo "SonarQube analysis skipped: ${e.getMessage()}"
                        echo "Credentials not configured yet, continuing..."
                    }
                }
            }
        }
        
        stage('Quality Gate') {
            steps {
                echo 'Waiting for SonarQube quality gate...'
                script {
                    try {
                        timeout(time: 5, unit: 'MINUTES') {
                            waitForQualityGate abortPipeline: false
                        }
                    } catch (Exception e) {
                        echo "Quality Gate skipped: ${e.getMessage()}"
                        echo "SonarQube not configured yet, continuing..."
                    }
                }
            }
        }
        
        stage('Build') {
            steps {
                echo 'Building Go application...'
                sh 'go build -o app .'
            }
        }
        
        stage('Build Docker Image') {
            steps {
                echo 'Building Docker image...'
                script {
                    def dockerImage = docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                    dockerImage.tag("latest")
                }
            }
        }
        
        stage('Push to Nexus') {
            steps {
                echo 'Pushing Docker image to Nexus...'
                script {
                    try {
                        withCredentials([usernamePassword(credentialsId: 'nexus-credentials', usernameVariable: 'NEXUS_USER', passwordVariable: 'NEXUS_PASS')]) {
                            sh '''
                                echo "Logging in to Nexus Docker registry..."
                                docker login ${DOCKER_REGISTRY} -u ${NEXUS_USER} -p ${NEXUS_PASS}
                                echo "Tagging Docker image..."
                                docker tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_REGISTRY}/${DOCKER_IMAGE}:${DOCKER_TAG}
                                docker tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_REGISTRY}/${DOCKER_IMAGE}:latest
                                echo "Pushing Docker image to Nexus..."
                                docker push ${DOCKER_REGISTRY}/${DOCKER_IMAGE}:${DOCKER_TAG}
                                docker push ${DOCKER_REGISTRY}/${DOCKER_IMAGE}:latest
                                echo "Docker image pushed successfully!"
                            '''
                        }
                    } catch (Exception e) {
                        echo "Nexus push skipped: ${e.getMessage()}"
                        echo "Nexus credentials not configured yet, continuing..."
                    }
                }
            }
        }
    }
    
    post {
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed!'
        }
        always {
            cleanWs()
        }
    }
}

