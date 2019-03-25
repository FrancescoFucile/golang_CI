node {
    def root = tool name: 'Go 1.8', type: 'go'
        stage('integrate'){
        sh 'mkdir bin'
        dir('src') {
            withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "GOPATH=/var/jenkins_home/workspace/proto_1/"]) {
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