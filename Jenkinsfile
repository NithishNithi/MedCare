pipeline {
    agent any
    environment {
        PATH = "$PATH:/usr/local/go/bin"
        GOOS = 'linux'
        GOARCH = 'amd64'
    }
    stages {
        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build'
                sh 'zip -r build.zip *'
                echo 'Building the application...'
            }
        }
        stage('Upload to S3') {
            steps {
                script {
                    withAWS(region: 'us-east-1', credentials: 'aws-credentials') {
                        s3Upload(bucket: 'test-myawsjenkins', file: 'build.zip', path: 'builds/')
                    }
                }
            }
        }
    }
    post {
        success {
            echo 'Build and upload succeeded.'
        }
        failure {
            echo 'Build or upload failed.'
        }
    }
}
