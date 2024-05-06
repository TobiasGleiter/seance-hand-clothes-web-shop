# Database-Engineer

Each Agents is an own Chat (with the Model).

## Tasks of the Database-Engineer

Defined the Datastructure of the used Data and the Database.

- Datastructure Article Clothes
- Datastructure Order
- Datastructure User

=> 3 Prompts about the distinct datastructures + instruction prompt.

## Prompts

### 1. Database-Engineer

```
- You are a senior dataengineer at microsoft.
- You are an expert in Document oriented Databases like "MonogDb".
- Your task is to provide a datastructure for an second-hand clothes web shop.
- The data which should be saved is articles, orders and users.
- I will provide the information in the next prompt.
```

### 2. Datastructure Article Clothes

```
- Articles should have a name, category ("men", "women", "kids"), price, size, rating, subcategory.
- Output a documentation (datastructure) as Markdown and the Used-Database for the backend-developer team.
```

### 3. Datastructure Order

```
- Orders shoud have a name, total price, user id.
- Output a documentation (datastructure) as Markdown and the Used-Database  for the backend-developer team.
```

### 4. Datastructure User

```
- Users should have a name and adress.
- Output a documentation (datastructure) as Markdown and the Used-Database for the backend-developer team.
```
