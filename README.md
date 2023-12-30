# Authentication Example in GO

This project showcases how to signup, login and authenticate a user with GO as a backend. 
The corresponding frontend flutter application can be found [here](https://github.com/activcoding/authenticaion-example-with-flutter).

## Requirements
Ensure you have the following prerequisites installed:
1. [GO](https://golang.org/)
2. [MongoDB](https://www.mongodb.com/)

And an account for the following services:

3. [MailJet](https://www.mailjet.com/)

## Setup
1. Clone the repository
2. Create a `.env` file in the root directory and add the following variables:
```
JWT_KEY=YourSigningKey
ISS=guardianGate://Random String
AUD=guardianKey://Random String
API_KEY=YourAPIKey
MAIL_JET_API_KEY=YourMailJetAPIKey
MAIL_JET_SECRET_KEY=YourMailJetSecretKey
```
Note: JWT_Key, ISS, AUD, and API_Key can be any random string. 
Ensure they match between the Flutter app and the backend. 

You can find your MailJet API Key and Secret Key in your MailJet account.

3. Run `go run main.go` to start the server on port 8081.
4. You can test the server with Postman or the [Flutter App](https://github.com/activcoding/authenticaion-example-with-flutter).

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://god.gw.postman.com/run-collection/28239926-594e326f-7098-45f2-a741-1085567a35d3?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D28239926-594e326f-7098-45f2-a741-1085567a35d3%26entityType%3Dcollection%26workspaceId%3D95dda62c-8934-47c8-91dd-9da44864034b)

## Available Routes
### Routes reachable without JWT
1. `/auth/signup`               - POST - Creates a new user
2. `/auth/signin`               - POST - Logs in a user and returns a JWT
3. `/auth/sendActivationEmail`  - POST - Verifies a users email address
4. `/auth/activateAccount`      - POST - Activates a user's account

### Routes reachable with JWT
1. `/exampleRestricted/exampleRestricted` - POST - Just a dummy route to test the authentication

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
