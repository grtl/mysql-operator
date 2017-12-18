// Poor people Enum
enum BuildStatus {
  STARTED,
  SUCCESS,
  FAILED
}

// Color defaults
COLOR_DEFAULT = "#dedede"
COLOR_SUCCESS = "#36a64f"
COLOR_FAILED  = "#f85050"

/**
 * Run given build stages and report status to slack
 * @param buildStages function containing all build stages
 **/
def slackNotifyStatus(buildStages) {
  try {
    def causes = currentBuild.rawBuild.getCauses().collect{
      it.getShortDescription()
    }.join(', ')
    slackSendStatus(BuildStatus.STARTED, "${causes}")
    buildStages()
    // Remove "and counting" suffix
    def buildTime = currentBuild.getDurationString()[0..-14]
    slackSendStatus(BuildStatus.SUCCESS, "${buildTime}")
  } catch (e) {
    slackSendStatus(BuildStatus.FAILED)
    throw e
  }
}

/**
 * Send slack message about build status
 * @param status build status (should be one of the BUILD_* constants)
 * @param extra additional information to be sent in a message
 **/
def slackSendStatus(BuildStatus status, String extra = "") {
  def messageColor
  def messageTitle
  def message = "${env.JOB_NAME} - #${env.BUILD_NUMBER}"

  switch (status) {
    case BuildStatus.STARTED:
      messageColor = COLOR_DEFAULT
      messageTitle = "${extra}"
      break;
    case BuildStatus.SUCCESS:
      messageTitle = "Success after ${extra}"
      messageColor = COLOR_SUCCESS
      break;
    case BuildStatus.FAILED:
      messageTitle = "Failed"
      messageColor = COLOR_FAILED
      break;
    default:
      messageTitle = "${extra}"
      messageColor = COLOR_DEFAULT
      break;
  }

  // Build message
  def messageBody = "${message} ${messageTitle} [<${env.RUN_DISPLAY_URL}|Open>]"
  slackSend color: "${messageColor}", message: "${messageBody}"
}

return this
