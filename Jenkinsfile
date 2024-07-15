pipeline {
    agent any

    stages {
        stage('Example') {
            steps {
                script {
                    withAWS(credentials: 'aws-credentials', region: 'us-west-2') {
                        // AWS-specific steps here
                        sh 'aws s3 ls'
                    }
                }
            }
        }
    }
}
