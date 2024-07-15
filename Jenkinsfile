pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'zip -r build.zip *'
                echo "Building the application..."
            }
        }
    }
    post {
        success {
            stage('Upload to S3') {
                steps {
                    withAWS(region: 'us-west-2', credentials: 'aws-credentials') {
                        s3Upload(bucket: 'jenkinstestbucket3', file: 'build.zip', path: 'builds/')
                    }
                }
            }
        }
        failure {
            echo 'Build or upload failed.'
        }
    }
}
