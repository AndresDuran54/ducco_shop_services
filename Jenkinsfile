pipeline {
    agent {
        label 'kubernetes'
    }

    stages {
        stage('Test') {
            steps {
                sh 'pwd'
                // sh 'docker build -t ducco_wallet:v2 ./@deploy/@micros/ducco_wallet'
            }
        }
    }
}
