node {
  def root = tool name: 'Go 1.9.2', type: 'go'

  ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/grtl/mysql-operator") {
    withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
      env.PATH="${GOPATH}/bin:$PATH"

      stage('Checkout') {
        checkout scm
      }

      stage('Install dependencies') {
        sh 'go get -u github.com/golang/dep/cmd/dep'
        sh 'dep ensure'
      }

      stage('Build') {
        sh 'go build'
      }
    }
  }
}
