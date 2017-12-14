node {
  def root = tool name: 'Go 1.9.2', type: 'go'
  def buildPath = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"

  ws("${buildPath}/src/github.com/grtl/mysql-operator") {
    withEnv(["GOROOT=${root}", "GOPATH=${buildPath}/", "PATH+GO=${root}/bin"]) {
      env.PATH="${GOPATH}/bin:$PATH"
      main()
    }
  }
}

def main() {
  // Checkout before loading additional groovy files
  stage('Checkout') {
    checkout scm
  }

  def slack = load "jenkins/slack.groovy"
  slack.slackNotifyStatus(this.&buildStages)
}

def buildStages() {
  stage('Install dependencies') {
    sh 'go get -u github.com/golang/dep/cmd/dep'
    sh 'dep ensure'
  }

  stage('Build') {
    sh 'go build'
  }
}
