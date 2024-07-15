pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh export PATH=$PATH:/usr/local/go/bin
                sh env GOOS=linux 
                sh env GOARCH=amd64
                sh go mod tidy
                sh go build 
                sh zip -r build.zip *
                sh 'echo "Building the application..."'
            }
        }
    }
    post {
        success {
            stage('Upload to S3') {
                steps {
                    withAWS(region: 'us-west-2', credentials: 'aws-credentials') {
                        s3Upload(bucket: 'jenkinstestbucket3', file: './build.zip', path: 'builds/')
                    }
                }
            }
        }
        failure {
            echo 'Build or upload failed.'
        }
    }
}
