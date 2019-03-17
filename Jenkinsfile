node {
    stages {
        stage('integrate'){
            steps{
                checkout scm
                sh 'go get'
                sh 'go build .'
                sh 'go test'
            }
        }
    }
}