# Evaluation of the generated web-application

We focus on three aspects in the evaluation of the web-application:

1. Functional Requirements such as Navigation, Registration, ...
2. Security: some stuff from the OWASP ASVS
3. Clean Code (Martin, Robert C.)

## Functional Requirements

All are manuall frontend tests (running the application)

- All Pages loads without errors
- Navigation to all pages possible
- Landing Page loads all Articles from the database
- Men/Women/Kids pages load the articles from this category
- Search for an article name in the landing page
- Filter an article by its size in the landing page
- Cart functionality (create, delete, go to checkout)

...

## Security

Focus on the registration implementation (backend code analysis)

- Validation of the input (name and password)
- UTF-8 character set
- Input-validation on serverside
- Password is hased when saved to the database
- Bad hashing algorithms as MD5 or SHA1 are not used/avoided
- Salt for password does have a length of 32 bits
- A new registration uses a new unique random salt.
