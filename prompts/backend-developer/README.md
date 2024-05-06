# Backend-Developer

Each Agents is an own Chat (with the Model).

## Tasks of the Backend-Developer

**(Group)**

- Fetch all articles from the database
- Fetch all articles with category women from the database
- Fetch all articles with category men from the database
- Fetch all articles with category kids from the database

**(Group)**

- Save an order to the database

**(Group)**

- Register a user and save to the database
- Login a user and save to the database

**(Group)**

- Fetch articles for a given user from the database
- Create an article for a given user and save to the database
- Delete an article for a given user from the database
- Update an article for a given user and save to the database

=> 4 Prompts + one instruction prompt

## Prompts

### 1. Backend-Developer

```
- You are a senior backend-developer at microsoft.
- You are an expert in "Go" programming and database interactions with "MongoDb"
- Your task is to implement the backend-funtionality for a second-hand clothes web shop.
- A databaseengineer provides you information about the datastructure.
- I will provide the more information in the next prompt.
```

### 2. Read articles

```
- Implement API-Endpoints with these functionalities:
1. Read all articles from the "MonogDB" database.
2. Read all articles from the "MongoDB" database that has the category women.
3. Read all articles from the "MongoDB" database that has the category men.
4. Read all articles from the "MongoDB" database that has the category kids.
- Output code and an API-Documentation that a frontend-developer can use.
- The Article Database structure is: {{DATABASEENGINEER_ARTCLES_STRUCTURE}}
```

### 3. Save order

```
- Implement an API-Endpoint with this functionality:
1. Retrieve order from the client and save it to the database.
- Output code and an API-Documentation that a frontend-developer can use.
- The Order Database structure is: {{DATABASEENGINEER_ORDER_STRUCTURE}}
```

### 4. Register and Login

```
- Implement API-Endpoints with these functionalities:
1. Register user and check if the user exits in the databasesave, if not, save to the datase.
2. Login user and validate if the user exists in the database.
- Output code and an API-Documentation that a frontend-developer can use.
- The Article Database structure is: {{DATABASEENGINEER_USER_STRUCTURE}}
```

### 5. CRUD articles

```
- Implement API-Endpoints with these functionalities:
1. Read articles for a given user from the database
2. Create an article for a given user and save to the database
3. Delete an article for a given user from the database
4. Update an article for a given user and save to the database
- Output code and an API-Documentation that a frontend-developer can use.
- The Article Database structure is: {{DATABASEENGINEER_USER_STRUCTURE}}
```
