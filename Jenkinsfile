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
        DOCKERHUB_CREDENTIALS=credentials('dockerhub')
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



    stage('Build') {

        steps {
            sh 'docker build -t lonesomet0wn/jsonplay:latest .'
        }
    }

    stage('Login') {

        steps {
            sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
        }
    }

    stage('Push') {

        steps {
            sh 'docker push lonesomet0wn/jsonplay:latest'
        }
    }
    }

    post {
        always {
            // Clean up our workspace.
            sh 'docker logout'
            deleteDir()
        }
    }
}   