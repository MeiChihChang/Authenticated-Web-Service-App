# The Authenticated Web App & Service

## Introduction

The App includes a web service where users authenticate through a login page and subsequently have the ability to download a dataset from the OpenData.Swiss platform after completing a simple task.

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
https://github.com/MeiChihChang/Authenticated-Web-Service-App/assets/37042542/a730c90f-7aef-4120-9f1c-f3fa9b5ed053
Follow the above workflow, you need to setup **Realm**, **Client**, **User**, **Roles**. And assign roles to the user then it will grant ths user's access permission to access opendata.swiss.

Set these Client_ID, Client_Secret, Realm name values from Keycloak at backend then this APP can adapt to your setting. 

## Setup servers

There are two servers for frontend and backend. Please download source codes then run the following commands:
```bash
frontend> npm install
frontend> npm start
backend> go run ./cmd/api .
```
Then appication will launch as follows:

https://github.com/MeiChihChang/Authenticated-Web-Service-App/assets/37042542/f5cc5a14-3a32-47c6-b878-6f2e421158d4

### How to use

Step 1. Login with your username & password as registerd at Keycloak
Step 2. Navigate to OpenData.Swiss to see the listed dataset from different organizations.
Step 3. You need to pass the math test then you can select the organization to see the available dataset downdload url.

### SWagger

There is swagger yaml file under backend/cmd/api/doc, which shows the server APIs
https://github.com/MeiChihChang/Authenticated-Web-Service-App/assets/37042542/830f3dec-8f73-4382-bbc2-6dacd313d350