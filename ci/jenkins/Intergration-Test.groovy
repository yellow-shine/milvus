@Library('jenkins-shared-library@v0.69.0') _

def pod = libraryResource 'io/milvus/pod/tekton-4am.yaml'
def milvus_helm_chart_version = '4.2.8'

pipeline {
    options {
        skipDefaultCheckout true
        parallelsAlwaysFailFast()
        buildDiscarder logRotator(artifactDaysToKeepStr: '30')
        preserveStashes(buildCount: 5)
        // abort previous build if it's a PR, otherwise queue the build
        disableConcurrentBuilds(abortPrevious: env.CHANGE_ID != null)
        timeout(time: 6, unit: 'HOURS')
        throttleJobProperty(
            categories: ['integration-test'],
            throttleEnabled: true,
            throttleOption: 'category'

        )
    }
    agent {
        kubernetes {
            cloud '4am'
            yaml pod
        }
    }
    stages {
        stage('meta') {
            steps {
                container('jnlp') {
                    script {
                        isPr = env.CHANGE_ID != null
                        gitMode = isPr ? 'merge' : 'fetch'
                        gitBaseRef = isPr ? "$env.CHANGE_TARGET" : "$env.BRANCH_NAME"
                    }
                }
            }
        }
        stage('build & test') {
            steps {
                container('tkn') {
                    script {
                        def job_name = tekton.integration_test arch: 'amd64',
                                              isPr: isPr,
                                              gitMode: gitMode ,
                                              gitBaseRef: gitBaseRef,
                                              pullRequestNumber: "$env.CHANGE_ID",
                                              make_cmd: "make clean && make build-cpp",
                                              test_entrypoint: "make integration-test",
                                              codecov_report_name: "integration-test",
                                              codecov_files: "./it_coverage.txt",
                                              tekton_log_timeout: '30m',
                                              tekton_pipeline_timeout: '3h'
                    }
                }
            }
            post {
                always {
                    container('tkn') {
                        script {
                            tekton.sure_stop()
                        }
                    }
                }
            }
        }
    }
}