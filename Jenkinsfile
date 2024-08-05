pipeline {
    agent {
        label 'kubernetes'
    }

    environment {
        deploy_service = 'ducco_products'
        deploy = 'dc-deploy-products'
        deploy_container = 'dc-products'
        img_v = 'v8'
    }

    stages {
        stage('Test') {
            steps {
                sh 'echo "Approved test!"'
            }
        }

        stage('Build') {
            steps {
                sh 'docker build -t andresduran54/$deploy_service:$img_v -f ./@deploy/@micros/$deploy_service/Dockerfile .'
                sh 'docker push andresduran54/$deploy_service'
                sh 'kubectl set image deployment/$deploy $deploy_container=$deploy_service:$img_v --record=true'
            }
        }
    }
}
