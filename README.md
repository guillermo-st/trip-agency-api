# trip-agency-api
A simple WebAPI with user authentication + authorization, implemented in Golang. This API's endpoints allow for:
* Authentication and Authorization of both admin users and the agency's drivers.
* Creation of new drivers (for admin users only).
* Listing all drivers with pagination, and wether they are currently travelling or not (for admin users only).
* Drivers can change their travelling status by starting and finishing trips.

#Assumptions
* The WebAPI will be consumed via HTTP containing JSON bodies. gRPC should be considered as more microservices join the system.
* An user can only have one role at a time (for now, either *admin* or *driver*). An user's permission level is stored as a custom claim in their JWT depending on their role, and will be safe from tampering due to the nature of JWT's signatures.
* New *admin* users can only be created by the DB administrator manually.

# Stack
[gin:]()
Gin was chosen initially as a Router for the serving API, but I quickly realized its benefits in terms of convenience when using **gin.Context**, **ShouldBindJSON**, and the many helpful abstractions it brings to routing and middleware. 

[sqlx:]()
A superset of the go sql package that allows unmarshalling query results into golang structs, making the process way less tedious.

[jwt/v4:]()
A very simple library for generating and validating JSON Web Tokens.

[pq:]()
The battle-tested PostgreSQL driver for golang. Postgres was my relational DB of choice, and a .sql script is included in this repo to recreate the tables.

# Known points of improvement
* Not everything is unit tested as it should, but one of the endpoints is unit tested all the way through as an example.
* Despite roles being persisted in the DB, there is a 1:1 representation of the role_id spectrum as constants in memory.
* Pagination is achieved with simple **LIMIT N OFFSET M** statements in the queries, and that won't scale performantly as the DB grows.

# Deployment
To deploy this service, I would take the following general steps:
1. Create a Heroku app for deployment of the API. This includes setting up the necessary environment variables and postgreSQL heroku plugins. For convenience, both of them could be in Docker containers.
2. Install Jenkins on a server and connect it with a GitHub webhook to push changes to it.
3. Configure Jenkins accordingly via ssh.
4. Instruct Jenkins to automatically deploy to Heroku on succesful testing of the master branch as a post-build action.
