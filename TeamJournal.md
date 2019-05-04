# Team - The Sharks

## Counter Burger Application

## Week 1
### Points Discussed
```
1. Projects ideas discussed
2. Each member searched for given SAAS application in the canvas.
```
### Outcome
```
1. We decided to build Counter Burger Online Order Application
2. Created initial version of Architecture diagram.
```
### Task Assigned
```
1. Read about building go apis.
2. Functionalities to work on like User, restaurants, payments etc.
3. Front end possibilities with Node.js and react.js
```

## Week 2
### Points Discussed
1. Backend technology - GO
2. Front end technology - php

### Action Points
1. Start with initial sample GO API on local
2. Setup mongo db environment
3. Go with mongo
3. Check out possibilities with php

### Blockers
```
Front end technology - Node.JS and React.JS
```

## Week 3
### Points discussed
1. Database schema for User, Restaurant, menu api
2. Tested User GO API with db for User
3. Front end pages development for login, signup, list restaurant

### Action Points
```
1. Create sample dataset for users, restaurant, and menu services.
2. Integrate with GO API
3. Work on front end pages
```

### Blocker
```
1. Routing between different APIs.
2. Maintaining state from Front end.
```

## Week 4
### Points Discussed
```
1. Blocker from last week 
2. Discussed APIs for Order, payment.
3. front end pages for cart, order and payment.
```
### Action Points
```
1. Implementation of microservices and other end points
2. Deployment of microservices on AWS
2. Application load balancer setup
3. front end add cart functionality
```
### Blockers
```
1. Add cart function
```

## Week 5
### Action Points
```
1. Integrate all services with front end
2. Test full application
3. Bug fixes
```
# AKF Scale Cube

### X-axis Scaling: 
- Horizontal duplication or x-axis scaling is to create multiple instances or clones of your application on AWS behind a load balancer.

- We have used auto scaling group based on CPU utilization and application load balancer..

### Y-axis Scaling:
- Y axis scaling or functional decomposition is to separate services or dissimilar things.

- We have implemented this by making separate microservices for users, restaurant, menu, order and payment.

### Z-axis Scaling:
- Z axis scaling is to decompose or splitting similar data into different chunks.

# Network Partition

Mongodb with 2 replicas are used for implementation. According to CAP theorem, MongoDB is CP database. Hence, it provides consistent data. Mongo DB is partition tolerant system, as it updates database from nwtowrk recovery and elects a new master node in that case.

# Architecture Diagram
![System Architecture](https://github.com/nguyensjsu/sp19-281-the-sharks/blob/master/Images/System%20Architecture.png)


# Application Screenshots





