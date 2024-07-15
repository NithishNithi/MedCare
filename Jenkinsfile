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
    }
    post {
        success {
            script {
                withAWS(credentials: 'aws-credentials', region: 'us-east-1') {
                    s3Upload(acl: 'Private', bucket: 'test-myawsjenkins', file: './build.zip')
                }
                echo 'Build and upload succeeded.'
            }
        }
        failure {
            echo 'Build or upload failed.'
        }
    }
}
