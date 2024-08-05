pipeline {
    agent {
        label 'kubernetes'
    }

    stages {
        stage('Test') {
            steps {
                sh 'docker build -t ducco_wallet:v1 ./@deploy/@micros/ducco_wallet'
            }
        }
    }
}
