# The Authenticated Web App & Service

## Introduction

This App includes a web service where users authenticate through a login page and subsequently have the ability to download a dataset from the OpenData.Swiss platform after completing a simple task.

Authenticated users are registered with Keycloak with specific role access permissions. They must be granted these role access permissions to have the right to view and download datasets from OpenData.Swiss.

This App aims to protect sensitive OpenData.Swiss dataset through:

- **Authenticity:** Only valid authenticated users with specific access permissions can log in and access datasets from OpenData.Swiss.

- **Confidentiality:** Sensitive datasets from OpenData.Swiss are only accessible to valid users.
  

## Prequirement

Install Go 1.8 or higher and Node.js v18.18.0 or higher, and Docker.


## Setup Keycloak

Keycloak is an open-source Identity and Access Management system installed via Docker using the following command:
```bash
docker run -d -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin --name keycloak jboss/keycloak:4.1.0.Final
```
![plot](https://github.com/MeiChihChang/Authenticated-Web-Service-App/assets/37042542/a730c90f-7aef-4120-9f1c-f3fa9b5ed053)
< Reference Image (Image source: https://medium.com/@allusaiprudhvi999/authentication-and-authorization-in-golang-microservice-using-an-open-source-iam-called-keycloak-46f03a26248f) >

Follow the above diagram to setup **Realm**, **Client**, **User**, **Roles**. Assign roles to users to grant access permissions to opendata.swiss.

Set the *Client_ID*, *Client_Secret*, and *Realm* values from Keycloak at the backend .env.local file so that this app can adapt to your settings.


## Setup servers

There are two servers for frontend and backend. Download the source code and run the following commands:
```bash
frontend> npm install
frontend> npm start
backend> go mod tidy
backend> go run ./cmd/api .
```
Then appication will launch as follows:

![plot](https://github.com/MeiChihChang/Authenticated-Web-Service-App/assets/37042542/f5cc5a14-3a32-47c6-b878-6f2e421158d4)

## How to use

Step 1. Log in with your username & password as registered at Keycloak.

Step 2. Navigate to OpenData.Swiss to see the listed datasets from different organizations.

Step 3. You need to pass the math test, then you can see this page.

Step 4. Select the organization to see the list of available dataset download URLs.


## Swagger

There is a Swagger YAML file under backend/cmd/api/doc, which shows the server APIs.

![plot](https://github.com/MeiChihChang/Authenticated-Web-Service-App/assets/37042542/830f3dec-8f73-4382-bbc2-6dacd313d350)


