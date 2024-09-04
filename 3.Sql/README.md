# Setting up PostgreSQL

**Use docker image of PostgreSQL to create a PostgreSQL server**

### To Set Up a PostgreSQL Server:

1. **Pull the Docker Image:**

   Run the following command in your terminal to pull the latest PostgreSQL image:

   ```bash
   docker pull postgres
   ```

2. **Start a PostgreSQL Instance:**

   Run the command below to create and start a PostgreSQL container:

   ```bash
   docker run --name <container_name> -e POSTGRES_PASSWORD=<password> -d postgres
   ```

   - `--name <container_name>`: Assigns a specific name to the container.
   - `-e POSTGRES_PASSWORD=<password>`: Sets the `POSTGRES_PASSWORD` environment variable inside the container.
   - `-d`: Runs the container in detached mode (in the background).
   - `postgres`: Specifies the PostgreSQL image to use.

   For example:

   ```bash
   docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
   ```

### Connecting to PostgreSQL:

**To Open Database Connection in GoLang:**

1. **Install the pq Package:**

   Use the following command to install the pq package, a pure Go Postgres driver for the `database/sql` package:

   ```bash
   go get github.com/lib/pq
   ```

2. **Connection String:**

   The connection string format is:

   ```plaintext
   postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable
   ```

   Example:

   ```plaintext
   postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable
   ```

   - `<username>`: PostgreSQL username (e.g., `postgres`).
   - `<password>`: PostgreSQL password.
   - `<host>`: Host where the PostgreSQL server is running (e.g., `localhost` or a container IP).
   - `<port>`: Port where PostgreSQL is listening (default is `5432`).
   - `<database>`: Name of the database you want to connect to (e.g., `postgres`).
   - `sslmode=disable`: Disables SSL encryption (common for local development).

**To Open a Database Connection:**

```go
func Open(driverName, dataSourceName string) (*DB, error)
```

- `driverName`: A string representing the name of the database driver (e.g., `"postgres"`).
- `dataSourceName`: A string representing the data source name, which typically contains the information needed to connect to the database (e.g., the username, password, database name, host, port).

**Returns:**

- `*DB`: A pointer to a `DB` type, representing a database handle used to interact with the database.
- `error`: An error type holding information if the connection fails.

### Executing Queries:

**To Run a Query:**

1. **Creating a Table:**

   Use `db.Exec` to execute a SQL statement such as creating a table:

   ```go
   func createProductTable(db *sql.DB) {
       query := `
           CREATE TABLE IF NOT EXISTS products (	
               id SERIAL PRIMARY KEY,
               name VARCHAR(50) NOT NULL,
               price NUMERIC(6, 2) NOT NULL,
               available BOOLEAN,
               created_at TIMESTAMP DEFAULT NOW()
           );`

       _, err := db.Exec(query)
       if err != nil {
           log.Fatal(err)
       }

       fmt.Println("Product table created successfully")
   }
   ```

2. **Inserting Data:**

   Use `db.Exec` with `RETURNING id` to insert data and retrieve the generated primary key:

   ```go
   func insertProduct(db *sql.DB, product Product) int {
       query := `INSERT INTO products (name, price, available)
           VALUES ($1, $2, $3) RETURNING id;`

       var pk int
       err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
       if err != nil {
           log.Fatal(err)
       }

       return pk
   }
   ```

3. **Querying Data by Primary Key:**

   Use `db.QueryRow` to query data by primary key:

   ```go
   func queryProductByID(db *sql.DB, id int) (Product, error) {
       query := `SELECT id, name, price, available, created_at FROM products WHERE id = $1;`

       var product Product
       err := db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Available, &product.CreatedAt)
       if err != nil {
           if err == sql.ErrNoRows {
               return product, nil
           }
           return product, err
       }

       return product, nil
   }
   ```

4. **Querying All Data:**

   Use `db.Query` to retrieve all records:

   ```go
   func queryAllProducts(db *sql.DB) ([]Product, error) {
       query := `SELECT id, name, price, available, created_at FROM products;`

       rows, err := db.Query(query)
       if err != nil {
           return nil, err
       }
       defer rows.Close()

       var products []Product
       for rows.Next() {
           var p Product
           err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Available, &p.CreatedAt)
           if err != nil {
               return nil, err
           }
           products = append(products, p)
       }

       if err := rows.Err(); err != nil {
           return nil, err
       }

       return products, nil
   }
   ```