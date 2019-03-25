node() {
  def root = tool name: 'Go 1.8', type: 'go'
  stage('Preparation') {
    checkout scm
  }
  stage ('test') {
   withEnv(["GOROOT=${root}", "GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin:${WORKSPACE}/bin", "GOBIN=${WORKSPACE}/bin"]){
    sh 'go get github.com/romana/rlog'
    sh 'go build'
    sh 'go test .'
   }
  }
}