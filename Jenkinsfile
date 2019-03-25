node {
    def root = tool name: 'Go 1.8', type: 'go'
    dir ('bin') {}
        stage('integrate'){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
            sh 'go version'
            checkout scm
            sh 'go get'
            sh 'go build .'
            sh 'go test'
        }
    }
}