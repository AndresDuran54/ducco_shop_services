pipeline {
    agent {
        label 'kubernetes'
    }

    environment {
        service = 'ducco_products'
    }

    stages {
        stage('Test') {
            steps {
                sh 'echo "Approved test"'
            }
        }

        stage('Build') {
            steps {
                sh 'echo $service'
                sh 'docker build -t andresduran54/$service -f ./@deploy/@micros/$service/Dockerfile .'
                sh 'docker push andresduran54/$service'
            }
        }
    }
}
