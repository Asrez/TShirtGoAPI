# T-Shirt Go API


# Endpoints

## API Endpoints

- `GET /categories`: Retrieve a list of available T-shirt categories.
- `GET /brands`: Retrieve a list of available T-shirt brands.
- `GET /sizes`: Retrieve a list of available T-shirt sizes.
- `GET /colors`: Retrieve a list of available T-shirt colors.
- `GET /materials`: Retrieve a list of available T-shirt materials.
- `GET /products`: Retrieve a list of available T-shirt products.

## API Examples
### Get T-shirt Categories:

Endpoint: GET /categories

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

Endpoint: GET /brands

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

Endpoint: GET /sizes

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

Endpoint: GET /colors

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

Endpoint: GET /materials

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

Endpoint: GET /products

Description: Retrieve a list of available T-shirt products.

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

Copyright 2023, Max Base
