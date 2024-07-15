pipeline{
    agent any

    stages{
        stage("Aws test credentials"){
            steps{
                withAWS(credentials: 'aws-credentials', region: 'us-east-1'){
                    sh 'echo "Hello DevOps" > hello.txt'
                    s3Upload (acl: 'Private' , bucket: 'test-myawsjenkins' , file: 'hello.txt')
                    sh "cat downloaded.txt"

                }
            }
        }
    }
}
