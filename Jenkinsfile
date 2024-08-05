pipeline {
    agent {
        label 'kubernetes'
    }

    environment {
        deploy_service = 'ducco_products'
    }

    stages {
        stage('Test') {
            steps {
                sh 'echo "Approved test!"'
            }
        }

        stage('Build') {
            steps {
                sh 'docker build -t andresduran54/$deploy_service:v6 -f ./@deploy/@micros/$deploy_service/Dockerfile .'
                sh 'docker push andresduran54/$deploy_service'
                sh 'kubectl set image deployment/dc-deploy-products dc-products=$deploy_service:v6'
            }
        }
    }
}
