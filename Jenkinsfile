pipeline {
    agent any

    stages {
        stage('Test') {
            steps {
                // Checkout the source code from the repository
                git 'https://github.com/AbdouTlili/testWebApp.git'
                
                // Run tests (replace with your actual test command)
                sh 'go test ./.'
            }
        }
        stage('Build Docker Image') {
            steps {
                // Build Docker image
                sh 'docker build -t abdoutlili/testwebapp:latest .'
            }
        }
    }

    post {
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
