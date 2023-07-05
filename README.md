# T-Shirt Go API

The T-Shirt Go API is a web-based application programming interface (API) designed to provide access to a wide range of T-shirt-related information and functionalities. The API enables developers to retrieve information about T-shirt categories, brands, sizes, colors, materials, and products.

## T-Shirt Key Features

- **T-Shirt Categories:** Retrieve a list of available T-shirt categories, such as men's T-shirts, women's T-shirts, and kids' T-shirts. This helps users navigate and filter T-shirt options effectively.
- **T-Shirt Brands:** Obtain a list of T-shirt brands, enabling users to identify their preferred brand options.
- **T-Shirt Sizes:** Retrieve a list of available T-shirt sizes, including small, medium, large, etc. This assists users in selecting the appropriate size for their desired T-shirt.
- **T-Shirt Colors:** Access a list of available T-shirt colors, allowing users to choose T-shirts that match their preferences.
- **T-Shirt Materials:** Retrieve information about different T-shirt materials, such as cotton, polyester, and blends. This helps users make informed decisions based on their preferred fabric.
- **T-Shirt Products:** Obtain a list of available T-shirt products, including details such as the T-shirt's name, category, brand, size, color, material, and price. Users can browse through the available options and make selections based on their preferences.

## API Endpoints

- `GET /categories`: Retrieve a list of available T-shirt categories.
- `GET /brands`: Retrieve a list of available T-shirt brands.
- `GET /sizes`: Retrieve a list of available T-shirt sizes.
- `GET /colors`: Retrieve a list of available T-shirt colors.
- `GET /materials`: Retrieve a list of available T-shirt materials.
- `GET /products`: Retrieve a list of available T-shirt products.

## API Examples
### Get T-shirt Categories:

Endpoint: `GET /categories`

Description: Retrieve a list of available T-shirt categories.

Example Response:

```json
[
  {
    "id": 1,
    "name": "Men's T-shirts"
  },
  {
    "id": 2,
    "name": "Women's T-shirts"
  },
  {
    "id": 3,
    "name": "Kids' T-shirts"
  }
]
```

### Get T-shirt Brands:

Endpoint: `GET /brands`

Description: Retrieve a list of available T-shirt brands.

Example Response:

```json
[
  {
    "id": 1,
    "name": "Nike"
  },
  {
    "id": 2,
    "name": "Adidas"
  },
  {
    "id": 3,
    "name": "Puma"
  }
]
```

### Get T-shirt Sizes:

Endpoint: `GET /sizes`

Description: Retrieve a list of available T-shirt sizes.

Example Response:

```json
[
  {
    "id": 1,
    "name": "Small"
  },
  {
    "id": 2,
    "name": "Medium"
  },
  {
    "id": 3,
    "name": "Large"
  }
]
```

### Get T-shirt Colors:

Endpoint: `GET /colors`

Description: Retrieve a list of available T-shirt colors.

Example Response:

```json
[
  {
    "id": 1,
    "name": "Red"
  },
  {
    "id": 2,
    "name": "Blue"
  },
  {
    "id": 3,
    "name": "Black"
  }
]
```

### Get T-shirt Materials:

Endpoint: `GET /materials`

Description: Retrieve a list of available T-shirt materials.

Example Response:

```json
[
  {
    "id": 1,
    "name": "Cotton"
  },
  {
    "id": 2,
    "name": "Polyester"
  },
  {
    "id": 3,
    "name": "Blend"
  }
]
```

### Get T-shirt Products:

Endpoint: `GET /products`

Description: Retrieve a list of available T-shirt products.

Endpoint: GET /products

Description: Retrieve a list of available T-shirt products, with optional filters for brand, color, size, and material.

Parameters:

- `brand (optional)`: Filter products by brand.
- `color (optional)`: Filter products by color.
- `size (optional)`: Filter products by size.
- `material (optional)`: Filter products by material.

Example Usage:

- `GET /products?brand=Nike`: Retrieve T-shirt products of the brand "Nike".
- `GET /products?color=Red&size=Large`: Retrieve T-shirt products that are red and large in size.
- `GET /products?brand=Adidas&color=Blue&size=Medium&material=Cotton`: Retrieve T-shirt products of the brand "Adidas" that are blue, medium-sized, and made of cotton.

Example Response:

```json
[
  {
    "id": 1,
    "name": "Men's Red T-shirt",
    "category": "Men's T-shirts",
    "brand": "Nike",
    "size": "Large",
    "color": "Red",
    "material": "Cotton",
    "price": 29.99
  },
  {
    "id": 2,
    "name": "Women's Blue T-shirt",
    "category": "Women's T-shirts",
    "brand": "Adidas",
    "size": "Medium",
    "color": "Blue",
    "material": "Polyester",
    "price": 24.99
  },
  {
    "id": 3,
    "name": "Kids' Black T-shirt",
    "category": "Kids' T-shirts",
    "brand": "Puma",
    "size": "Small",
    "color": "Black",
    "material": "Blend",
    "price": 19.99
  }
]
```

## Database PostgreSQL

```sql
-- Table: Categories
CREATE TABLE Categories (
  category_id INT PRIMARY KEY AUTO_INCREMENT,
  category_name VARCHAR(255)
);

-- Table: Brands
CREATE TABLE Brands (
  brand_id INT PRIMARY KEY AUTO_INCREMENT,
  brand_name VARCHAR(255)
);

-- Table: Sizes
CREATE TABLE Sizes (
  size_id INT PRIMARY KEY AUTO_INCREMENT,
  size_name VARCHAR(50)
);

-- Table: Colors
CREATE TABLE Colors (
  color_id INT PRIMARY KEY AUTO_INCREMENT,
  color_name VARCHAR(50)
);

-- Table: Materials
CREATE TABLE Materials (
  material_id INT PRIMARY KEY AUTO_INCREMENT,
  material_name VARCHAR(50)
);

-- Table: Products
CREATE TABLE Products (
  product_id INT PRIMARY KEY AUTO_INCREMENT,
  product_name VARCHAR(255),
  category_id INT,
  brand_id INT,
  size_id INT,
  color_id INT,
  material_id INT,
  price DECIMAL(10, 2),
  FOREIGN KEY (category_id) REFERENCES Categories(category_id),
  FOREIGN KEY (brand_id) REFERENCES Brands(brand_id),
  FOREIGN KEY (size_id) REFERENCES Sizes(size_id),
  FOREIGN KEY (color_id) REFERENCES Colors(color_id),
  FOREIGN KEY (material_id) REFERENCES Materials(material_id)
);
```

Copyright 2023, Max Base
