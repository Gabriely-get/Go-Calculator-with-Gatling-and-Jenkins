pipeline {
    agent any
    tools { go 'go-1.20' }

    stages {
        stage('Run_Application') {
            steps {
                sh 'chmod 744 /var/jenkins_home/workspace/Calculator-Gatling-Stress-Test/app/'
                sh '/var/jenkins_home/workspace/Calculator-Gatling-Stress-Test/app/ go build .'
                sh '/var/jenkins_home/workspace/Calculator-Gatling-Stress-Test/app/ go run .'
            }
        }
        stage('Download_Gatling') {
            steps {
                sh "curl -LO https://repo1.maven.org/maven2/io/gatling/highcharts/gatling-charts-highcharts-bundle/3.9.0/gatling-charts-highcharts-bundle-3.9.0-bundle.zip"
            }
        }
    }
}