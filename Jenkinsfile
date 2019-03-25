node {
    def root = tool name: 'Go 1.8', type: 'go'
    stage('integrate'){
        withEnv(["GOROOT=${root}", "GOPATH=${root}/bin"]) {
            sh 'go version'
            checkout scm
            sh 'go get'
            sh 'go build .'
            sh 'go test'
        }
    }
}