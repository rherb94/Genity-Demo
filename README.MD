# Ryan Herb GO RESTful API with Docker 

Here contains the GO project code for a RESTful api demo project.  

**Instructions:**
- After cloning repository, you will have to build the Docker image: `docker build -t genity-demo .`

- after the image is build, you can start the container: `docker run -p 5000:5000 genity-demo`

**Testing With Postman**
- Once the container is running, the api can be tested by submitting a POST request against localhost:5000/post-data/Name and you will see a "Post Successful" message in the response body. 

- Next, we can check if this succeeded by submitting a GET request against localhost:5000/get-data and you will see the newly persisted Data object containing the path variable that was passed in the POST request as well as the generated UUID4 and TimeStamp values



Project Report PDF: [Recipe-ManagerReport](/Recipe-ManagerReport.pdf?raw-true)

![POST](/assets/GET-Message.png?raw=true "POST")