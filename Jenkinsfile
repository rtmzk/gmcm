#!groovy
pipeline {
    agent {
        label 'golang-ci'
    }

    options {
        timestamps()
    }

    environment {
        GOOS = 'linux'
        GOARCH = 'amd64'
    }

    parameters {
        booleanParam(name: 'binary', defaultValue: true, description: '构建二进制')
        booleanParam(name: 'rpm', defaultValue: false ,description: '构建rpm包')
        booleanParam(name: 'corsPlatform', defaultValue: false, description: '跨平台构建')
    }

    stages {
        stage('clone gmcm repository') {
            steps {
                dir("${env.WORKSPACE}") {
                    sh 'git config --global http.sslVerify false'
                    git branch: 'master',
                    credentialsId: '01066ccb-3845-4fc6-95ef-c122aeda1ffc',
                    url: 'http://192.168.254.25:40001/ops/gmcm.git'
                }
            }
        }
        stage ('build binary') {
            when {
                expression { params.binary }
                not {
                    expression { params.corsPlatform }
                }
            }
            steps {
                dir("${env.WORKSPACE}") {
                    make build.binary
                }
            }
        }
        stage ('build rpm') {
            when {
                expression { params.rpm }
            }
            steps {
                dir("${env.WORKSPACE}") {
                //TODO...
                }
            }
        }

        stage ('cors platform build') {
            when {
                allOf {
                    expression { params.binary }
                    expression { params.corsPlatform }
                }
            }
            steps {
                dir("${env.WORKSPACE}") {
                    make build.multiarch
                }
            }
        }

        stage ('copy binary to 77') {
            when {
                expression { params.binary }

                not {
                    expression { params.corsPlatform }
                }
            }
            steps {
                dir("${env.WORKSPACE}"){
                    scp
                }
            }
        }
        stage ('copy rpm to 77') {
            when {
                expression { params.rpm }
            }
            steps {
                dir("${env.WORKSPACE}") {
                    //TODO...
                }
            }
        }
        stage ('copy all platform binary to 77') {
            when {
                expression { params.corsPlatform }
            }
            steps {

            }
        }
    }
}