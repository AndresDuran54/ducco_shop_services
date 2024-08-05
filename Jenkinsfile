pipeline {
    agent {
        label 'kubernetes'
    }

    stages {
        stage('Test') {
            steps {
                sh 'pwd'
                sh 'tree'
                sh 'docker build -t andresduran54/ducco_wallet:v5 -f ./@deploy/@micros/ducco_wallet/Dockerfile .'
                sh 'docker push andresduran54/ducco_wallet:v5'
            }
        }
    }
}
