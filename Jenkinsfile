node {
    def root = tool name: 'Go 1.8', type: 'go'
        stage('integrate'){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}, GOBIN=${GOPATH}/bin"]) {
            sh 'go version'
            checkout scm
            sh 'go get'
            sh 'go build .'
            sh 'go test'
        }
    }
}