node {
    stage('integrate'){
        checkout scm
        sh 'go get'
        sh 'go build .'
        sh 'go test'
    }
}