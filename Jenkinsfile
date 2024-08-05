pipeline {
    agent {
        label 'kubernetes'
    }

    stages {
        stage('Test') {
            steps {
                sh 'docker ps -a'
            }
        }
    }
}
