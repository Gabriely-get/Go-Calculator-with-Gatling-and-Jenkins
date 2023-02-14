# Calculator with Gatling Stress Test

## Required

- Install and configuring Docker
- Download this project

## Configuring Jenkins

1. Pull Jenkins docker image:
   -       docker pull jenkins/jenkins:lts
2. Run the Jenkins container:
   -     docker run --name jenkins --rm -p 8282:8080 -p 50000:50000 -p 8000:8000 -v jenkins_home:/var/jenkins_home -v /var/run/docker.sock:/var/run/docker.sock -v $(which docker):$(which docker) jenkins/jenkins:lts
3. Access Jenkins: http://localhost:8282/
   ##### The password will be shown in the terminal
4. Install suggested plugins, click on *Skip and continue as admin* or *Create First Admin User* then Verify the URL, finish and start
5. On *Manage Jenkins*, go to *Manage Plugins* and install on *Available plugins* the *Go* Plugin for GoLang
6. Install the *Gatling* plugin too
7. Still in *Manage Jenkins* go to *Global Tool Configuration*. In *Add Go*, set it to Install automatically from golang.org, name it as *go-1.20* apply and save it.

## Running with Jenkins

### Creating jobs
1. On *Dashboard* click in `New Item`
2. Set the name, choose `Pipeline` and click `OK`
   #### Name job 1 as *Calculator* and job 2 as *Gatling Stress Test*
3. Go to `Pipeline` and choose *Pipeline script from SCM* and fill as following:
- SCM: Git 
- Repository URL: https://github.com/Gabriely-get/Go-Calculator-with-Gatling-and-Jenkins.git
- Branch Specifier: */main
- Choose the Jenkinsfile path. The steps to create the two jobs are the same, the only thing that will change is the path:
   1. ./job1/Jenkinsfile
   2. ./job2/Jenkinsfile
6. Run first Calculator job and after Gatling Stress Test job
   ##### ps: Calcualtor job only stops if you abort it manually

### Getting Results
1. When the job *Gatling Stress Test* finish, click in the finished build and go to *Workspaces* of the job
2. Click on the folder link
3. Follow the path: *gatling-charts-highcharts-bundle-3.9.0/results/*
4. Choose the result and click on *all files in zip*
5. Extract it and open the *index.html* with a Browser
6. The result will look like the image below: 
![gatling-index](/images/gatling-index.png)
