pipeline {
    agent {
        label 'kubernetes'
    }

    stages {
        stage('Test') {
            steps {
                sh 'docker start $(docker ps -qa)'
            }
        }
    }
}
