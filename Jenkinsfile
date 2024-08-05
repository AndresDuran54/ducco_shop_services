pipeline {
    agent {
        label 'kubernetes'
    }

    stages {
        stage('Test') {
            steps {
                sh 'echo "$(docker ps -a)"'
            }
        }
    }
}
