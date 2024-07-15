pipeline {
    agent any

    stages {
        stage("Aws test credentials") {
            steps {
                script {
                    // Define AWS credentials and region
                    withAWS(credentials: 'aws-credentials', region: 'us-east-1') {
                        // Run commands within AWS context
                        sh 'echo "Hello DevOps" > hello.txt'
                        // Upload file to S3 bucket
                        s3Upload(file: 'hello.txt', bucket: 'test-myawsjenkins')
                        // Example of downloading a file (if needed)
                        sh "aws s3 cp s3://test-myawsjenkins/hello.txt downloaded.txt"
                    }
                }
            }
        }
    }
}
