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
4. Install suggested plugins, click on *Skip and continue as admin* then Verify the URL, finish and start
5. On *Manage Jenkins*, go to *Plugin Manager* and install the *Go Plugin* for GoLang
6. Still in *Manage Jenkins* go to *Global Tool Configuration*. In *Add Go*, set it to Install automatically from golang.org, name it as *go-1.20* and save it.

## Running with Jenkins

### Creating jobs
1. On *Dashboard* click in `New Item`
2. Set the name, choose `Pipeline` and click `OK`
   #### Choose appropriated names as: *Calculator* for job 01 and *Gatling Stress Test* for job 02.
3. Go to `Pipeline` and choose *Pipeline script from SCM* and fill as following:
- `SCM`: Git 
- Repository URL: https://github.com/Gabriely-get/Go-Calculator-with-Gatling-and-Jenkins.git
- Branch Specifier: */main
- Choose the Jenkinsfile path. The steps to create the two jobs are the same, the only thing that will change is the path:
   1. ./job1/Jenkinsfile
   2. ./job2/Jenkinsfile
6. Run first Calculator job and after Gatling Stress Test job
   ##### ps: Calcualtor job only stops if you abort it manually

## Running without Jenkins

### Calculator
1. Open the *src* folder project on terminal
2. Type: *go build .*
3.  Type: *go run .*

### Gatling
1. Open the extracted Gatling folder on terminal
2. Type: *bin/recorder.sh*

## Endpoints
Endpoints and examples of appropriate response

- Addition: http://localhost:8000/calc/sum/{value1}/{value2}
   ``` json
     {
        "result": "10"
     }
     ``` 
- Subtraction: http://localhost:8000/calc/sub/{value1}/{value2}
  ``` json
     {
        "result": "5"
     }
     ``` 
- Division: http://localhost:8000/calc/div/{value1}/{value2}
  ``` json
     {
        "result": "7"
     }
     ``` 
- Multiplication: http://localhost:8000/calc/mult/{value1}/{value2}
  ``` json
     {
        "result": "30"
     }
     ``` 
- History: http://localhost:8000/calc/history
  ``` json
    [
        {
            "time": "2022-08-19T02:13:52.853500749-03:00",
            "result": "6.00 - 6.00 = 0.00",
            "operation": "SUBTRACTION"
        },
        {
            "time": "2022-08-19T02:13:52.859714718-03:00",
            "result": "6.00 + 6.00 = 12.00",
            "operation": "ADDITION"
        }
    ]
     ``` 
## Testing

1. Open the folder project on terminal
2. Type: *sudo chmod +x ./smoke-test-calculator.sh*
3. Type: *./smoke-test-calculator.sh* to execute the test