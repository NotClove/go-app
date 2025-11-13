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
                dir('test/go-app') {
                    echo 'Using local project directory'
                }
            }
        }
        
        stage('Test') {
            steps {
                echo 'Running Go tests...'
                dir('test/go-app') {
                    sh 'go test -v -coverprofile=coverage.out ./...'
                }
            }
            post {
                always {
                    dir('test/go-app') {
                        sh 'go tool cover -html=coverage.out -o coverage.html || true'
                        publishHTML([
                            reportName: 'Coverage Report',
                            reportDir: 'test/go-app',
                            reportFiles: 'coverage.html',
                            keepAll: true
                        ])
                    }
                }
            }
        }
        
        stage('SonarQube Analysis') {
            steps {
                echo 'Running SonarQube analysis...'
                withSonarQubeEnv('SonarQube') {
                    dir('test/go-app') {
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
                }
            }
        }
        
        stage('Quality Gate') {
            steps {
                echo 'Waiting for SonarQube quality gate...'
                timeout(time: 5, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: false
                }
            }
        }
        
        stage('Build') {
            steps {
                echo 'Building Go application...'
                dir('test/go-app') {
                    sh 'go build -o app .'
                }
            }
        }
        
        stage('Build Docker Image') {
            steps {
                echo 'Building Docker image...'
                dir('test/go-app') {
                    script {
                        def dockerImage = docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                        dockerImage.tag("latest")
                    }
                }
            }
        }
        
        stage('Push to Nexus') {
            steps {
                echo 'Pushing Docker image to Nexus...'
                script {
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

