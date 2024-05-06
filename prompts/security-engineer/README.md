# Security-Engineer

Each Agents is an own Chat (with the Model).

## Tasks

Only for the Register and Login pages.

- Validate the incoming data (Length, Datatype, Structure)
- Enforce a character set (UTF-8)
- Enforce data validation on serverside
- Validate if password is hashed and salted before saved to database.

=> 2 Prompts for Register and Login pages

## Prompts

### 1. Security-Engineer

```
- You are a senior security engineer at microsoft.
- You are an expert making secure web-applications, especially backends.
- Your task is to check code if its secure.
- I will provide the information in the next prompt.
```

### 2. Secure Register Functionality

```
- Update the code to enforce this security features:
1. Validate the incoming data (Length, Datatype, Structure)
2. Enforce a character set (UTF-8)
3. Enforce data validation on serverside
4. Validate if password is hashed and salted before saved to database.
- Output updated, secure code.
```

### 3. Secure Login Functionality

```
- Update the code to enforce this security features:
1. Validate the incoming data (Length, Datatype, Structure)
2. Enforce a character set (UTF-8)
3. Enforce data validation on serverside
- Output updated, secure code.
```
