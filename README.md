I couldnt make it to convert this project into using docker containers.
To start this project, first create an .env file. And fill up the actual values as per given by the .env.sample file.
Make sure mariadb is running.
Next run the command: `CompileDaemon -command="./cognixus_todo_api"`.
Before calling any api, user has to login first.
User can log in with either Google or Github.
They will be verified by their session token.
