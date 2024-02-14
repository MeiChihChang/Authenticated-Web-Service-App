# The Authenticated Web App & Service

## Introduction

This App includes a web service where users authenticate through a login page and subsequently have the ability to download a dataset from the OpenData.Swiss platform after completing a simple task.

The authenticated users are registered at Keycloak with specific role access permission. They must be granted this role access permission then they have the right to see and download dataset from OpenData.Swiss.

This App aims to protect sensitive OpenData.Swiss dataset through:

- **Authenticity:** Only valid authenticated users with spefic access permission can login and access dataset from OpenData.Swiss.

- **Confidentiality:** Senstive dataset from OpenData.Swiss to download is only limited valid users.
  

## Prequirement

Install Go 1.8 higher and node.js v18.18.0 higher and docker


## Setup Keycloak

Keycloak is an Open Source Identity and Access Management and install by docker as following command:
```bash
docker run -d -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin --name keycloak jboss/keycloak:4.1.0.Final
```
![plot](https://github.com/MeiChihChang/Authenticated-Web-Service-App/assets/37042542/a730c90f-7aef-4120-9f1c-f3fa9b5ed053)
< The above image is from https://medium.com/@allusaiprudhvi999/authentication-and-authorization-in-golang-microservice-using-an-open-source-iam-called-keycloak-46f03a26248f >

Follow the above diagram, you need to setup **Realm**, **Client**, **User**, **Roles**. And assign roles to the user then it will grant ths user's access permission to access opendata.swiss.

Set these Client_ID, Client_Secret, Realm values from Keycloak at backend *.env.local* file then this APP can adapt to your setting. 


## Setup servers

There are two servers for frontend and backend. Please download source codes then run the following commands:
```bash
frontend> npm install
frontend> npm start
backend> go mod tidy
backend> go run ./cmd/api .
```
Then appication will launch as follows:

![plot](https://github.com/MeiChihChang/Authenticated-Web-Service-App/assets/37042542/f5cc5a14-3a32-47c6-b878-6f2e421158d4)

## How to use

Step 1. Login with your username & password as registerd at Keycloak

Step 2. Navigate to OpenData.Swiss to see the listed dataset from different organizations.

Step 3. You need to pass the math test then you can see this page.

Step 4. Select the organization to see the list of available dataset downdload urls.


## SWagger

There is a swagger yaml file under backend/cmd/api/doc, which shows the server APIs

![plot](https://github.com/MeiChihChang/Authenticated-Web-Service-App/assets/37042542/830f3dec-8f73-4382-bbc2-6dacd313d350)
