pipeline {
    agent any

    stages {
        stage("File Creation") {
            steps {
                // Create a file in Jenkins workspace
                sh 'echo "Hello DevOps" > hello.txt'
            }
        }
        
        stage("Upload to S3") {
            steps {
                // Upload file to S3 bucket
                script {
                    withAWS(credentials: 'aws-credentials', region: 'us-east-1') {
                        s3Upload(file: 'hello.txt', bucket: 'test-myawsjenkins')
                    }
                }
            }
        }
    }
}
