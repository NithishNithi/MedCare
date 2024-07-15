pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'zip -r build.zip *'
                echo "Building the application..."
            }
        }
        stage('Upload to S3') {
            steps {
                script {
                    def awsCredentials = 'aws-credentials' // Update with your credentials ID
                    def bucketName = 'jenkinstestbucket3'
                    def filePath = 'build.zip'
                    def s3Path = 'builds/'

                    withAWS(region: 'us-west-2', credentials: awsCredentials) {
                        awsS3Upload(file: filePath, bucket: bucketName, path: s3Path)
                    }
                }
            }
        }
    }
    post {
        success {
            echo 'Build and upload succeeded!'
        }
        failure {
            echo 'Build or upload failed.'
        }
    }
}
