node() {
  def root = tool name: 'Go 1.8', type: 'go'
  stage('Preparation') {
    checkout scm
  }
  stage ('Compile') {
    sh "${root}/bin/go build"
  }
  stage ('Test') {
   withEnv(["GOROOT=${root}", GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin:${WORKSPACE}/bin", "GOBIN=${WORKSPACE}/bin"]){
    sh 'go build'
    sh 'go get'
    sh 'go test .'
   }
  }
  stage ('Archive') {
    archiveArtifacts '**/tests.out, **/tests.xml, **/coverage.out, **/coverage.xml, **/coverage2.xml'
  }
}