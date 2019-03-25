node {
    def root = tool name: 'Go 1.8', type: 'go'
        stage('integrate'){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}"]) {
            sh 'ls'
            sh 'go version'
            checkout scm
            sh 'go get'
            sh 'go build .'
            sh 'go test'
        }
    }
}