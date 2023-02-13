pipeline {
    agent any

    stages {
        stage('Run_Application') {
            sh 'go build .'
            sh 'go run .'
        }
        stage('Download_Gatling') {
            steps {
                sh "curl -LO https://repo1.maven.org/maven2/io/gatling/highcharts/gatling-charts-highcharts-bundle/3.9.0/gatling-charts-highcharts-bundle-3.9.0-bundle.zip"
            }
        }
    }
}