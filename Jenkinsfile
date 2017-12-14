node {
  def root = tool name: 'Go 1.9.2', type: 'go'

  try {
    notifyStarted()

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

    notifySuccess()
  } catch (e) {
    notifyFailed()
    throw e
  }
}

def notifyStarted() {
  slackSend(message: "STARTED: Job '${env.JOB_NAME} #${env.BUILD_NUMBER}' (<${env.BUILD_URL}|Open>)")
}

def notifySuccess() {
  slackSend(color: '#36a64f', message: "SUCCESS: Job '${env.JOB_NAME} #${env.BUILD_NUMBER}' (<${env.BUILD_URL}|Open>)")
}

def notifyFailed() {
  slackSend(color: '#f85050', message: "FAILED: Job '${env.JOB_NAME} #${env.BUILD_NUMBER}' (<${env.BUILD_URL}|Open>)")
}
