// Poor people Enum
BUILD_STARTED = 0
BUILD_SUCCESS = 1
BUILD_FAILED  = 2

// Color defaults
COLOR_DEFAULT = "#dedede"
COLOR_SUCCESS = "#36a64f"
COLOR_FAILED  = "#f85050"

/** Run given build stages and report status to slack
 * @param buildStages function containing all build stages
 **/
def slackNotifyStatus(buildStages) {
  try {
    slackSendStatus(BUILD_STARTED)
    buildStages()
    // Remove "and counting" suffix
    def buildTime = currentBuild.getDurationString()[0..-14]
    slackSendStatus(BUILD_SUCCESS, "${buildTime}")
  } catch (e) {
    slackSendStatus(BUILD_FAILED)
    throw e
  }
}

/** Send slack message about build status
 * @param status build status (should be one of the BUILD_* constants)
 * @param extra additional information to be sent in a message
 **/
def slackSendStatus(status = BUILD_STARTED, String extra = "") {
  def messageColor
  def messageTitle
  def message = "${env.JOB_NAME} - #${env.BUILD_NUMBER}"

  switch (status) {
    case BUILD_STARTED:
      messageColor = COLOR_DEFAULT
      messageTitle = "Started"
      break;
    case BUILD_SUCCESS:
      messageTitle = "Success after ${extra}"
      messageColor = COLOR_SUCCESS
      break;
    case BUILD_FAILED:
      messageTitle = "Failed"
      messageColor = COLOR_FAILED
      break;
    default:
      messageTitle = "${extra}"
      messageColor = COLOR_DEFAULT
      break;
  }

  // Build message
  def messageBody = "${message} ${messageTitle} (<${env.BUILD_URL}|Open>)"
  slackSend color: "${messageColor}", message: "${messageBody}"
}

return this
