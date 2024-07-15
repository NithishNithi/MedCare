pipeline {
    agent any

    environment {
        AWS_DEFAULT_REGION = 'us-east-1'
    }

    stages {
        stage("Aws test credentials") {
            steps {
                withAWS(region: 'us-east-1', credentials: 'aws-credentials') {
                    // Run commands within AWS context
                    sh 'echo "Hello DevOps" > hello.txt'
                    // Upload file to S3 bucket
                    s3Upload(bucket: 'test-myawsjenkins', file: 'hello.txt')
                    // Example of downloading a file (if needed)
                    sh "aws s3 cp s3://test-myawsjenkins/hello.txt downloaded.txt"
                }
            }
        }
    }
}
