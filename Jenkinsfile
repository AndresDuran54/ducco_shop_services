pipeline {
    agent {
        label 'kubernetes'
    }

    stages {
        stage('Test') {
            steps {
                sh 'echo "Approved test"'
            }
        }

        stage('Build') {
            steps {
                sh '''docker build
                    -t andresduran54/ducco_wallet -f
                    ./@deploy/@micros/ducco_wallet/Dockerfile . > /dev/null'''
                sh 'docker push andresduran54/ducco_wallet'
            }
        }
    }
}
