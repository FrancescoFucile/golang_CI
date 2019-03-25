node {
    def root = tool name: 'Go 1.8', type: 'go'
        stage('integrate'){
        dir('src') {
            withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "GOPATH=${JENKINS_HOME}/workspace/${JOB_NAME}"]) {
                sh 'ls'
                sh 'go version'
                checkout scm
                sh 'go get'
                sh 'go build .'
                sh 'go test'
            }
        }
    }
}