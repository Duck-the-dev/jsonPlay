#!/usr/bin/env groovy
// The above line is used to trigger correct syntax highlighting.

pipeline {
    // Lets Jenkins use Docker for us later.
    agent any
    tools {
        go 'go1.19'
    }
    environment {
        GO119MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    // If anything fails, the whole Pipeline stops.
    stages {
        stage('Build & Test') {
            // Use golang.

            steps {
                // Build the app.
                sh 'go build'
            }
        }


        stage('Docker') {
            environment {
                // Extract the username and password of our credentials into "DOCKER_CREDENTIALS_USR" and "DOCKER_CREDENTIALS_PSW".
                // (NOTE 1: DOCKER_CREDENTIALS will be set to "your_username:your_password".)
                // The new variables will always be YOUR_VARIABLE_NAME + _USR and _PSW.
                // (NOTE 2: You can't print credentials in the pipeline for security reasons.)
                DOCKER_CREDENTIALS = credentials('dockerhub')
            }

            steps {
                // Use a scripted pipeline.
                script {
                    node {
                        def app

                        stage('Clone repository') {
                            checkout scm
                        }

                        stage('Build image') {
                            app = docker.build("${env.DOCKER_CREDENTIALS_USR}/jsonplay")
                        }

                        stage('Push image') {
                            // Use the Credential ID of the Docker Hub Credentials we added to Jenkins.
                            docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                                // Push image and tag it with our build number for versioning purposes.
                                app.push("${env.BUILD_NUMBER}")

                                // Push the same image and tag it as the latest version (appears at the top of our version list).
                                app.push("latest")
                            }
                        }
                    }
                }
            }
        }
    }

    post {
        always {
            // Clean up our workspace.
            deleteDir()
        }
    }
}   