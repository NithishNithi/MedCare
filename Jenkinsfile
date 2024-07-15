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
                // Upload file to S3 bucket using AWS CLI
                withCredentials([[$class: 'AmazonWebServicesCredentialsBinding', credentialsId: 'aws-credentials']]) {
                    sh 'aws s3 cp hello.txt s3://test-myawsjenkins/'
                }
            }
        }
    }
}
