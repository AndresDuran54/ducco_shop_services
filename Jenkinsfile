pipeline {
    agent {
        label 'kubernetes'
    }

    stages {
        stage('Test') {
            steps {
                sh 'pwd'
                sh 'tree'
                sh 'docker build -t ducco_wallet:v2 -f ./@deploy/@micros/ducco_wallet/Dockerfile .'
            }
        }
    }
}
