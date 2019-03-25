node {
    def root = tool name: 'Go 1.8', type: 'go'
    stage('integrate'){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "GOPATH=${WORKSPACE}"]) {
            sh 'go get'
            sh 'go build .'
            sh 'go test'
        }
    }
}